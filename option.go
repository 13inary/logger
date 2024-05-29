package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newOption 额外的选项
func newOption() []zap.Option {
	return []zap.Option{
		zap.AddCaller(), zap.AddCallerSkip(1), // 输出调用日志所在代码位置、跳过多少层调用栈
		zap.AddStacktrace(zapcore.ErrorLevel), // 输出整个调用栈
		//zap.Fields(zap.String("serverName", "api")), // 日志共有信息
	}
}
