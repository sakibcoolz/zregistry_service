package log

import (
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Log initialized")

	return logger
}
