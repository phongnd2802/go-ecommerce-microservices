package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"
)

type TaskDistributor interface {
	DistributeTaskSendOTPEmail(
		ctx context.Context,
		payload *PayloadSendOTPEmail,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}


func NewRedisTaskDistributor(redisSetting settings.RedisSetting) TaskDistributor {
	redisOtp := asynq.RedisClientOpt{
		Addr: redisSetting.Addr(),
	}
	client := asynq.NewClient(redisOtp)
	return &RedisTaskDistributor{
		client: client,
	}
}
