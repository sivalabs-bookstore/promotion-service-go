package application

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func makeLoggerForDevelopment() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func makeLoggerForProduction() (*zap.Logger, error) {
	prodConfig := zap.NewProductionConfig()
	prodConfig.Encoding = "console"
	prodConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	prodConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	return prodConfig.Build()
}
