package services

import (
	"context"

	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
)

type (
	UserAuth interface {
		Register(ctx context.Context, req *dto.RegisterRequest) (*repo.UserUserVerify, error)
		VerifyOTP(ctx context.Context, req *dto.VerifyOTPRequest) (*repo.UserUserVerify, error)
		UpdatePasswordRegister(ctx context.Context, req *dto.SetPasswordRequest) (*repo.UserUserProfile, error)
	}
)
