package api

import (
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	config  *user.Config
	ua services.UserAuth
}

func NewServer(config *user.Config, ua services.UserAuth) *Server {
	return &Server{
		config: config,
		ua: ua,
	}
}