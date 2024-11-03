package main

import (
	"fmt"
	"net"
	"os"

	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/app"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/config"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	DEVELOPMENT = "development"
	PRODUCTION  = "production"
)

func main() {
	cfg := &UserConfig{}
	err := config.LoadConfig("./configs", "user", "yaml", cfg)
	if err != nil {
		log.Fatal().Msgf("failed to read config %v", err)
	}
	if cfg.Grpc.Mode == DEVELOPMENT {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	go runTaskProcessor(cfg)
	runGrpcServer(cfg)

}

func runTaskProcessor(cfg *UserConfig) {
	taskProcessor, err := app.InitTaskProcessor(cfg.Redis, cfg.Email)
	if err != nil {
		log.Fatal().Msgf("cannot create task Proccessor: %v", err)
	}
	log.Info().Msgf("Start task processor")
	err = taskProcessor.Start()
	if err != nil {
		log.Fatal().Msgf("failed to start task processor")
	}
}

func runGrpcServer(cfg *UserConfig) {
	grpcLogger := grpc.UnaryInterceptor(logger.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	server, err := app.InitServer(cfg.DB, cfg.Redis)
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	if err != nil {
		log.Fatal().Msgf("cannot start gRPC server: %v", err)
	}
	network := "tcp"
	address := fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port)

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal().Msgf("cannot create listener")
	}

	log.Info().Msgf("Start GRPC server at: %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msgf("cannot start gRPC server: %v", err)
	}

}
