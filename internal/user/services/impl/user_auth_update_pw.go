package impl

import (
	"context"

	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/errs"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/postgres"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/crypto"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/random"
)

// UpdatePasswordRegister implements services.UserAuth.
func (ua *userAuthImpl) UpdatePasswordRegister(
	ctx context.Context, 
	req *dto.SetPasswordRequest,
) (*repo.UserUserProfile, error) {
	// Check Info Verify
	infoVerify, err := ua.store.GetUserVerifyByKeyHash(ctx, req.GetToken())
	if err != nil {
		if err == postgres.ErrRecordNotFound {
			return nil, errs.NotFoundError("token invalid")
		}
		return nil, errs.InternalError("failed to get info verify: %s", err)
	}
	if !infoVerify.IsVerified.Bool {
		return nil, errs.UnauthenticatedError("email has not been verified")
	}

	// Random Salt
	ranSalt := random.RandomString(10)
	// Hash Password
	hashedPassword, err := crypto.HashPasswordWithSalt(req.GetPassword(), ranSalt)
	if err != nil {
		return nil, errs.InternalError("failed to hash password: %s", err)
	}

	// Execute Transaction Set Password
	arg := repo.UpdatePassswordParamsTx{
		CreateUserBaseParams: repo.CreateUserBaseParams{
			UserEmail: infoVerify.VerifyKey,
			UserPassword: hashedPassword,
			UserSalt: ranSalt,
		},
	}
	result, err := ua.store.UpdatePasswordRegisterTx(ctx, arg)
	if err != nil {
		return nil, errs.InternalError("failed to create user profile: %s", err)
	}

	return &result.UserProfile, nil
}