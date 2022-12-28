package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Logger *zap.SugaredLogger

//func NewZapLogger() *zap.SugaredLogger {
//	cfg := zap.Config{
//		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
//		Encoding:    "json",
//		OutputPaths: []string{"stderr"},
//		EncoderConfig: zapcore.EncoderConfig{
//			MessageKey:   "message",
//			LevelKey:     "level",
//			TimeKey:      "time",
//			EncodeTime:   CustomTimeEncoder,                // format hiển thị thời gian log
//			EncodeCaller: zapcore.FullCallerEncoder,        // lấy dòng code bắt đầu log
//			EncodeLevel:  zapcore.CapitalColorLevelEncoder, // format cách hiển thị level
//			CallerKey:    "caller",
//		},
//	}
//
//	logger, _ := cfg.Build() // build logger from config
//	return logger.Sugar()
//}

func NewZapLogger() *zap.SugaredLogger {
	writer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	logger := zap.New(core)
	return logger.Sugar()
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./app.log")
	return zapcore.AddSync(file)
}
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		EncodeTime:   CustomTimeEncoder,           // format hiển thị thời gian log
		EncodeCaller: zapcore.ShortCallerEncoder,  // lấy dòng code bắt đầu log
		EncodeLevel:  zapcore.CapitalLevelEncoder, // format cách hiển thị level
		CallerKey:    "caller",
	})
}
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
