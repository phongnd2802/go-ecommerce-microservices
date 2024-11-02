package services

import (
	"context"

	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
)

type (
	UserRegister interface {
		Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
		VerifyOTP(ctx context.Context, req *dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error)
		UpdatePasswordRegister(ctx context.Context, req *dto.SetPasswordRequest) (*dto.SetPasswordResponse, error)
	}
)
