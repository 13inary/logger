package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newFileJsonEncoder 日志为json格式
func newFileJsonEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// newConsoleTextEncoder 日志为控制台文本格式
func newConsoleTextEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = customTimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// newFileTextEncoder 日志为文件文本格式
func newFileTextEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder // 日志输出调用的完整路径
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// customTimeEncoder 自定义日志的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
