package logger

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type WorkerLogger struct{}

func NewWorkerLogger() *WorkerLogger {
	return &WorkerLogger{}
}

func (logger *WorkerLogger) Print(level zerolog.Level, args ...interface{}) {
	log.WithLevel(level).Msg(fmt.Sprint(args...))
}

func (logger *WorkerLogger) Debug(args ...interface{}) {
	logger.Print(zerolog.DebugLevel, args...)
}

func (logger *WorkerLogger) Info(args ...interface{}) {
	logger.Print(zerolog.InfoLevel, args...)
}

func (logger *WorkerLogger) Warn(args ...interface{}) {
	logger.Print(zerolog.WarnLevel, args...)
}

func (logger *WorkerLogger) Error(args ...interface{}) {
	logger.Print(zerolog.ErrorLevel, args...)
}

func (logger *WorkerLogger) Fatal(args ...interface{}) {
	logger.Print(zerolog.FatalLevel, args...)
}