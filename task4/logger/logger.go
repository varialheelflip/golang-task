package logger

import "go.uber.org/zap"

var LOGGER *zap.Logger

func InitLogger() {
	LOGGER, _ = zap.NewProduction()
	defer LOGGER.Sync()
}
