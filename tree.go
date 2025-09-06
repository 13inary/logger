package logger

import "go.uber.org/zap/zapcore"

// newTree 多个core
func newTree(debugModel bool) zapcore.Core {
	cores := []zapcore.Core{
		newConsoleCore(debugModel),
		newFileCore(),
	}
	return zapcore.NewTee(cores...) // 可以设置多个core
}
