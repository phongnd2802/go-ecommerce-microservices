package impl

import (
	"context"

	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
)

// UpdatePasswordRegister implements services.UserAuth.
func (ur *userAuthImpl) UpdatePasswordRegister(ctx context.Context, req *dto.SetPasswordRequest) (*dto.SetPasswordResponse, error) {
	panic("unimplemented")
}