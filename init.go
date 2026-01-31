package logger

import (
	"go.uber.org/zap"
)

var (
	fasterLogger *zap.Logger
	superLogger  *zap.SugaredLogger
)

// InitLogger 初始化日志模块。日志使用方法1：这里管理日志模块
//
//	logPath 若为空，则表示使用默认路径：./log/service.log
func InitLogger(debugModel bool, logPath string) error {
	fasterLogger, superLogger = newLogger(debugModel, logPath, 1)
	return nil
}

// NewLogger 初始化日志模块。日志使用方法2：引用的项目自己管理日志模块
//
//	logPath 若为空，则表示使用默认路径：./log/service.log
//	callerSkip 调用栈跳过层数
func NewLogger(debugModel bool, logPath string) (*zap.Logger, *zap.SugaredLogger) {
	return newLogger(debugModel, logPath, 0)
}

func newLogger(debugModel bool, logPath string, callerSkip int) (*zap.Logger, *zap.SugaredLogger) {
	if logPath != "" {
		logFilePath = logPath
	}

	core := newTree(debugModel)
	options := newOption(callerSkip)
	fLogger := zap.New(core, options...)
	sLogger := fLogger.Sugar()
	return fLogger, sLogger
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
