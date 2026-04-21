package logger

import "go.uber.org/zap/zapcore"

// newTree 多个core
func newTree(consoleDebugModel bool, fileDebugModel bool) zapcore.Core {
	cores := []zapcore.Core{
		newConsoleCore(consoleDebugModel),
		newFileCore(fileDebugModel),
	}
	return zapcore.NewTee(cores...) // 可以设置多个core
}
