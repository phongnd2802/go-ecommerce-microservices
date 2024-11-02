package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const TaskSendOTPEmail = "task:send_otp_email"

type PayloadSendOTPEmail struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}


// DistributeTaskSendOTPEmail implements TaskDistributor.
func (distributor *RedisTaskDistributor) DistributeTaskSendOTPEmail(
	ctx context.Context, 
	payload *PayloadSendOTPEmail, 
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}	
	task := asynq.NewTask(TaskSendOTPEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}
	log.Println(info.Queue)
	return nil
}

// ProcessTaskSendOTPEmail implements TaskProcessor.
func (processor *RedisTaskProcessor) ProcessTaskSendOTPEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendOTPEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	subject := "Welcome to NDP-Ecommerce"
	content := fmt.Sprintf(`Hello,<br/> 
	Thank you for registering with us!<br/>
	OTP is %s<br/>`, payload.OTP)
	to := []string{payload.Email}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send otp email: %w", err)
	}

	return nil
}
