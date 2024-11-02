package api

import (
	"context"

	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
	dto "github.com/phongnd2802/go-ecommerce-microservices/pb/user"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	config  *user.Config
	ur services.UserRegister
}

func (server *Server) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	return server.ur.Register(ctx, req)
}

func (server *Server) VerifyOTP(ctx context.Context, req *dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, error) {
	return server.ur.VerifyOTP(ctx, req)
}

func (server *Server) UpdatePasswordRegister(ctx context.Context, req *dto.SetPasswordRequest) (*dto.SetPasswordResponse, error) {
	return server.ur.UpdatePasswordRegister(ctx, req)
}

func NewServer(config *user.Config, ur services.UserRegister) *Server {
	return &Server{
		config: config,
		ur: ur,
	}
}