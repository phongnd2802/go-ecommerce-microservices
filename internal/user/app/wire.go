//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services/impl"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/worker"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/cache"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/email"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/postgres"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"
)


func InitServer(postgresSetting settings.PostgresSetting, redisSetting settings.RedisSetting) (*Server, error) {
	panic(wire.Build(
		NewServer,
		impl.NewUserAuth,
		repo.NewStore,
		cache.NewRedisCache,
		worker.NewRedisTaskDistributor,
		newDBEngine,
	))
}

func InitTaskProcessor(
	redisSettings settings.RedisSetting,
	emailSetting settings.EmailSetting,
) (worker.TaskProcessor, error) {
	panic(wire.Build(
		worker.NewRedisTaskProcessor,
		email.NewGmailSender,
	))
}

func newDBEngine(cfg settings.PostgresSetting) (*pgxpool.Pool, error) {
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		return nil, err
	}
	return db.GetDB(), nil
}
