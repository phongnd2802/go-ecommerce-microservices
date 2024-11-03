package impl

import (
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/repo"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/services"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/worker"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/cache"
)


type userAuthImpl struct {
	cache cache.Cache
	store repo.Store
	taskDistributor worker.TaskDistributor
}


func NewUserAuth(store repo.Store, cache cache.Cache, taskDistributor worker.TaskDistributor) services.UserAuth {
	return &userAuthImpl{
		store: store,
		cache: cache,
		taskDistributor: taskDistributor,
	}
}
