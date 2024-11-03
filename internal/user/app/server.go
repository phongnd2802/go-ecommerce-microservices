package app

import (
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	config  *user.Config
	ua services.UserAuth

}

func NewServer(config *user.Config, ua services.UserAuth, grpcServer *grpc.Server) *Server {
	server := Server{
		config: config,
		ua: ua,
	}

	pb.RegisterUserServiceServer(grpcServer, &server)
	reflection.Register(grpcServer)
	return &server
}
