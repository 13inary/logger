package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newOption 额外的选项
func newOption(callerSkip int) []zap.Option {
	options := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel), // 输出整个调用栈
		//zap.Fields(zap.String("serverName", "api")), // 日志共有信息
	}

	// 输出调用日志所在代码位置、跳过多少层调用栈
	if callerSkip > 0 {
		options = append(options, zap.AddCallerSkip(callerSkip))
	}

	return options
}
