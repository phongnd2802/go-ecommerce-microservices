package impl

import (
	"context"
	"strconv"
	"time"

	"github.com/hibiken/asynq"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/worker"
	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/cache"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/postgres"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/crypto"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/random"
	"github.com/rs/zerolog/log"
)


const (
	TIME_OTP_REGISTERED = 3 // Minute
)


// Register implements services.UserAuth
func (ur *userAuthImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*repo.UserUserVerify, error) {
	// Hash Email
	hashedEmail := crypto.GetHash(req.GetVerifyKey())
	
	log.Debug().Str("hashEmail", hashedEmail)

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
		log.Info().Msg("Key does not exist")
	case err != nil:
		return nil, errs.InternalError("Get failed::", err)
	case otpFound != "":
		return nil, errs.ConflictError("")
	}

	// Generate OTP
	otpNew := random.GenerateSixDigit()
	log.Debug().Int("OTP", otpNew)

	//Execture Register Transaction
	// Save OTP to Postgres
	arg := repo.RegisterTxParams{
		CreateUserVerifyParams: repo.CreateUserVerifyParams{
			VerifyOtp:  strconv.Itoa(otpNew),
			VerifyKey:     req.GetVerifyKey(),
			VerifyKeyHash: hashedEmail,
		},
		AfterCreate: []func(repo.UserUserVerify) error {
			// Save OTP in Redis with expiration time
			func(uv repo.UserUserVerify) error {
				err = ur.cache.Set(ctx, userKey, uv.VerifyOtp, time.Duration(TIME_OTP_REGISTERED)*time.Minute)
				if err != nil {
					return errs.InternalError("failed to set otp to redis: %s", err)
				}
				return nil
			},
			// Send Email
			func(uv repo.UserUserVerify) error {
				taskPayload := &worker.PayloadSendOTPEmail{
					Email: uv.VerifyKey,
					OTP:   uv.VerifyOtp,
				}
				opts := []asynq.Option{
					asynq.MaxRetry(10),
					asynq.ProcessIn(10 * time.Second),
					asynq.Queue(worker.QueueCritical),
				}
				
				return ur.taskDistributor.DistributeTaskSendOTPEmail(ctx, taskPayload, opts...)
			},
		},
	}
	txResult, err := ur.store.RegisterTx(ctx, arg)
	if err != nil {
		if postgres.ErrorCode(err) == postgres.UniqueViolation {
			return nil, errs.ConflictError("email already exists: %s", err)
		}
		return nil, errs.InternalError("failed to register: %s", err)
	}
	
	return &txResult.UserVerify, nil
}
