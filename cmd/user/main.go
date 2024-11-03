package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hibiken/asynq"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/api"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services/impl"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/worker"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/cache"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/config"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/email"
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
	cache := cache.NewRedisCache(fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port))


	redisOpt := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	ur := impl.NewUserAuth(store, cache, taskDistributor)
	mailer := email.NewGmailSender(cfg.Email.EmailSenderName, cfg.Email.EmailSenderAddress, cfg.Email.EmailSenderPassword)

	
	go runTaskProcessor(redisOpt, mailer)
	runGrpcServer(cfg, ur)
}


func runTaskProcessor(redisOpt asynq.RedisClientOpt, mailer email.EmailSender) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, mailer)
	log.Println("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal("failed to start task processor")
	}
}

func runGrpcServer(cfg *user.Config, ua services.UserAuth) {
	server := api.NewServer(cfg, ua)

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