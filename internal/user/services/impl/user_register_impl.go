package impl

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hibiken/asynq"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/worker"
	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/cache"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/crypto"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/random"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	TIME_OTP_REGISTERED = 3 // Minute
)


type userRegisterImpl struct {
	cache cache.Cache
	store repo.Store
	taskDistributor worker.TaskDistributor
}

// Register implements services.UserRegister.
func (ur *userRegisterImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	violations := validateRegisterRequest(req)
	if violations != nil {
		return nil, errs.InvalidArgumentError(violations)
	}

	// Hash Email
	hashedEmail := crypto.GetHash(req.GetVerifyKey())
	fmt.Println(hashedEmail)
	
	// Check user exists in user base table
	userFound, err := ur.store.CheckUserBaseExists(ctx, req.GetVerifyKey())
	if err != nil {
		return nil, errs.InternalError("failed to check user base exists %s", err)
	}

	if userFound > 0 {
		return nil, errs.ConflictError("email already exists")
	}

	// Check OTP
	userKey := user.GetUserKeyOTP(hashedEmail)
	otpFound, err := ur.cache.Get(ctx, userKey)
	
	switch {
	case err == cache.ErrKeyNotFound:
		log.Println("Key does not exist")
	case err != nil:
		return nil, errs.InternalError("Get failed::", err)
	case otpFound != "":
		return nil, errs.ConflictError("")
	}

	// Generate OTP
	otpNew := random.GenerateSixDigit()
	log.Printf("OTP is :: %d\n", otpNew)

	// Save OTP to Postgres
	userVerify, err := ur.store.CreateUserVerify(ctx, repo.CreateUserVerifyParams{
		VerifyOtp: strconv.Itoa(otpNew),
		VerifyKey: req.GetVerifyKey(),
		VerifyKeyHash: hashedEmail,
	})

	if err != nil {
		return nil, errs.InternalError("failed to create user verify: %s", err)
	}

	// Save OTP in Redis with expiration time
	err = ur.cache.Set(ctx, userKey, strconv.Itoa(otpNew), time.Duration(TIME_OTP_REGISTERED)*time.Minute)
	if err != nil {
		return nil, errs.InternalError("failed to set otp to redis: %s", err)
	}

	// Send Email
	taskPayload := &worker.PayloadSendOTPEmail{
		Email: userVerify.VerifyKey,
		OTP: strconv.Itoa(otpNew),
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = ur.taskDistributor.DistributeTaskSendOTPEmail(ctx, taskPayload, opts...)
	if err != nil {
		return nil, errs.InternalError("%w", err)
	}
	rsp := &dto.RegisterResponse{
		VerifyId: userVerify.VerifyID,
	}
	return rsp, nil
}

// UpdatePasswordRegister implements services.UserRegister.
func (ur *userRegisterImpl) UpdatePasswordRegister(ctx context.Context, req *dto.SetPasswordRequest) (*dto.SetPasswordResponse, error) {
	panic("unimplemented")
}

// VerifyOTP implements services.UserRegister.
func (ur *userRegisterImpl) VerifyOTP(ctx context.Context, req *dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error) {
	panic("unimplemented")
}

func NewUserRegister(store repo.Store, cache cache.Cache, taskDistributor worker.TaskDistributor) services.UserRegister {
	return &userRegisterImpl{
		store: store,
		cache: cache,
		taskDistributor: taskDistributor,
	}
}

func validateRegisterRequest(req *dto.RegisterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateEmail(req.GetVerifyKey()); err != nil {
		violations = append(violations, errs.FieldViolation("verify_key", err))
	}
	return violations
}
