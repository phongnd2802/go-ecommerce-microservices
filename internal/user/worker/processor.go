package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/phongnd2802/go-ecommerce-microservices/internal/user/email"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendOTPEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	mailer email.EmailSender
}

// Start implements TaskProcessor.
func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendOTPEmail, processor.ProcessTaskSendOTPEmail)

	return processor.server.Start(mux)
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, mailer email.EmailSender) TaskProcessor {
	server := asynq.NewServer(redisOpt, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
		},
	})

	return &RedisTaskProcessor{
		server: server,
		mailer: mailer,
	}
}
