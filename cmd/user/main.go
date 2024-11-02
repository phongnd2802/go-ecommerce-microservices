package main

import (
	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/api"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services/impl"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main() {
	cfg := &user.Config{}
	err := config.LoadConfig("./configs", "user", "yaml", cfg)
	if err != nil {
		log.Fatalf("failed to read config %v", err)
	}

	cfgDB := cfg.DB
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
	 		cfgDB.Username, cfgDB.Password, cfgDB.Host, cfgDB.Port, cfgDB.DbName, cfgDB.SslMode)
	connPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := repo.NewStore(connPool)
	ur := impl.NewUserRegister(store)	
	server := api.NewServer(cfg, ur)


	gRPCServer := grpc.NewServer()
	pb.RegisterUserServiceServer(gRPCServer, server)
	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port))
	if err != nil {
		log.Fatalf("cannot create listener %v", err)
	}
	log.Println("start gRPC user service: ", listener.Addr().String())
	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server %v", err)
	}
}