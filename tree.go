package logger

import "go.uber.org/zap/zapcore"

// newTree 多个core
func newTree() zapcore.Core {
	cores := []zapcore.Core{
		newConsoleCore(),
		newFileCore(),
	}
	return zapcore.NewTee(cores...) // 可以设置多个core
}
