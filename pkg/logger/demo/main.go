package main

import (
	"go.uber.org/zap"
	"go01-airbnb/pkg/logger"
)

func main() {
	// Khởi tạo zap
	sugarLogger := logger.NewZapLogger()
	sugarLogger.Infof("Hello world", zap.String("OK", "Oke"))
}
