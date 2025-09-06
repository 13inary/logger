package logger

import (
	"go.uber.org/zap"
)

var (
	fasterLogger *zap.Logger
	superLogger  *zap.SugaredLogger
)

// InitLogger 初始化日志模块
// logPath 若为空，则表示使用默认路径：./log/service.log
func InitLogger(logPath string) error {
	if logPath != "" {
		logFilePath = logPath
	}

	core := newTree()
	options := newOption()
	fasterLogger = zap.New(core, options...)
	superLogger = fasterLogger.Sugar()
	return nil
}

// Sync 同步缓存
func Sync() error {
	if err := fasterLogger.Sync(); err != nil {
		return err
	}
	if err := superLogger.Sync(); err != nil {
		return err
	}
	return nil
}
