package app

import (
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	ua services.UserAuth
}

func NewServer(ua services.UserAuth) *Server {
	server := Server{
		ua: ua,
	}
	return &server
}
