package impl

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/cache"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/crypto"
)

// VerifyOTP implements services.UserAuth.
func (ua *userAuthImpl) VerifyOTP(ctx context.Context, req *dto.VerifyOTPRequest) (*repo.UserUserVerify, error) {
	hashedEmail := crypto.GetHash(req.GetVerifyKey())

	// Get OTP in Redis
	otpFound, err := ua.cache.Get(ctx, user.GetUserKeyOTP(hashedEmail))
	if err != nil {
		if err == cache.ErrKeyNotFound {
			return nil, errs.NotFoundError("invalid OTP or OTP has expired: %s", err)
		}
		return nil, errs.InternalError("failed to get otp: %s", err)
	}

	if req.GetVerifyCode() != otpFound {
		return nil, errs.NotFoundError("invalid OTP or OTP has expired")
	}

	// Get User Verify
	userVerify, err := ua.store.GetUserVerifyByKeyHash(ctx, hashedEmail)
	if err != nil {
		return nil, errs.InternalError("failed to get user verify: %s", err)
	}

	// Update Verified 
	result, err := ua.store.UpdateUserVerify(ctx, repo.UpdateUserVerifyParams{
		IsVerified: pgtype.Bool{
			Bool: true,
			Valid: true,
		},
		VerifyKeyHash: userVerify.VerifyKeyHash,
	})

	if err != nil {
		return nil, errs.InternalError("failed to update verify: %s", err)
	}

	return &result, nil
}