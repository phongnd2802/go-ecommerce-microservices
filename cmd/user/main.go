package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hibiken/asynq"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/app"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/config"
	"google.golang.org/grpc"
)

func main() {
	cfg := &user.Config{}
	err := config.LoadConfig("./configs", "user", "yaml", cfg)
	if err != nil {
		log.Fatalf("failed to read config %v", err)
	}

	go runTaskProcessor(cfg)
	runGrpcServer(cfg)

}

func runTaskProcessor(cfg *user.Config) {
	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.Redis.Addr(),
	}
	taskProcessor, err := app.InitTaskProcessor(redisOpt, cfg.Email)
	if err != nil {
		log.Fatalf("cannot create task Proccessor: %v", err)
	}
	log.Println("start task processor")
	err = taskProcessor.Start()
	if err != nil {
		log.Fatal("failed to start task processor")
	}
}


func runGrpcServer(cfg *user.Config) {
	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.Redis.Addr(),
	}
	grpcServer := grpc.NewServer()
	_, err := app.InitServer(cfg, cfg.DB, cfg.Redis, redisOpt, grpcServer)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %v", err)
	}
	network := "tcp"
	address := fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port)

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("cannot create listener")
	}

	log.Printf("start GRPC server at: %s\n", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %v", err)
	}

}
