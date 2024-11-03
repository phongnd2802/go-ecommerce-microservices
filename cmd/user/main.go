package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/app"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	cfg := &user.Config{}
	err := config.LoadConfig("./configs", "user", "yaml", cfg)
	if err != nil {
		log.Fatalf("failed to read config %v", err)
	}

	go runTaskProcessor(cfg)
	go runGatewayServer(cfg)
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
	server, err := app.InitServer(cfg, cfg.DB, cfg.Redis, redisOpt)
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
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

func runGatewayServer(cfg *user.Config) {
	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.Redis.Addr(),
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	server, err := app.InitServer(cfg, cfg.DB, cfg.Redis, redisOpt)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}
	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterUserServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatalf("cannot register handler server: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	network := "tcp"
	address := fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port)
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("cannot create listener: %v", err)
	}

	log.Printf("start HTTP Gateway server: %s\n", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatalf("cannot start HTTP Gateway server: %v", err)
	}
}
