package impl

import (
	"context"

	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
)

// VerifyOTP implements services.UserAuth.
func (ur *userAuthImpl) VerifyOTP(ctx context.Context, req *dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error) {
	panic("unimplemented")
}