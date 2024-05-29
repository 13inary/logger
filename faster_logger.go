package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// log method

func Debug(msg string, args ...zapcore.Field) {
	fasterLogger.Debug(msg, args...)
}

func Info(msg string, args ...zapcore.Field) {
	fasterLogger.Info(msg, args...)
}

func Warn(msg string, args ...zapcore.Field) {
	fasterLogger.Warn(msg, args...)
}

func Error(msg string, args ...zapcore.Field) {
	fasterLogger.Error(msg, args...)
}

func Fatal(msg string, args ...zapcore.Field) {
	fasterLogger.Fatal(msg, args...)
}

func Panic(msg string, args ...zapcore.Field) {
	fasterLogger.Panic(msg, args...)
}

// log key value

func Str(key string, val string) zapcore.Field {
	return zap.String(key, val)
}

func Int(key string, val int) zapcore.Field {
	return zap.Int(key, val)
}

func Int64(key string, val int64) zapcore.Field {
	return zap.Int64(key, val)
}

func Float64(key string, val float64) zapcore.Field {
	return zap.Float64(key, val)
}

func Bool(key string, val bool) zapcore.Field {
	return zap.Bool(key, val)
}

func Err(err error) zapcore.Field {
	return zap.Error(err)
}

func Time(key string, val time.Time) zapcore.Field {
	return zap.Time(key, val)
}

func Duration(key string, val time.Duration) zapcore.Field {
	return zap.Duration(key, val)
}

//func Object(key string, val zapcore.ObjectEncoder) zapcore.Field {
//	return zap.Object(key, val)
//}

func Any(key string, val any) zapcore.Field {
	return zap.Any(key, val)
}
