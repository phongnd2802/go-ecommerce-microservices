package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/email"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/logger"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/settings"
	"github.com/rs/zerolog/log"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	Stop()
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


// Stop implements TaskProcessor.
func (processor *RedisTaskProcessor) Stop() {
	if processor.server != nil {
		processor.server.Stop()
	}
}

func NewRedisTaskProcessor(redisSetting settings.RedisSetting, mailer email.EmailSender) TaskProcessor {
	redisOpt := asynq.RedisClientOpt{
		Addr: redisSetting.Addr(),
	}
	server := asynq.NewServer(redisOpt, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Str("type", task.Type()). 
				Bytes("payload", task.Payload()).
				Msg("process task failed")
		}),
		Logger: logger.NewWorkerLogger(),
	})

	return &RedisTaskProcessor{
		server: server,
		mailer: mailer,
	}
}
