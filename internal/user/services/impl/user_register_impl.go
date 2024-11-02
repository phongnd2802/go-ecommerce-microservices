package impl

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/crypto"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/random"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type userRegisterImpl struct {
	store repo.Store
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
	
	rsp := &dto.RegisterResponse{}
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

func NewUserRegister(store repo.Store) services.UserRegister {
	return &userRegisterImpl{
		store: store,
	}
}

func validateRegisterRequest(req *dto.RegisterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateEmail(req.GetVerifyKey()); err != nil {
		violations = append(violations, errs.FieldViolation("verify_key", err))
	}
	return violations
}
