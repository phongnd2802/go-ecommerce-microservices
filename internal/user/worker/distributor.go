package worker

import (
	"context"

	"github.com/hibiken/asynq"
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


func NewRedisTaskDistributor(redisOtp asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOtp)
	return &RedisTaskDistributor{
		client: client,
	}
}
