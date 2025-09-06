package logger

import (
	"go.uber.org/zap/zapcore"
)

// newConsoleCore 控制台输出的方式
func newConsoleCore(debugModel bool) zapcore.Core {
	encoder := newConsoleTextEncoder()

	consoleWriteSyncer := newConsoleWriter()
	writerSyncer := zapcore.NewMultiWriteSyncer(consoleWriteSyncer) // 可以指定多个输出

	logLevel := zapcore.InfoLevel
	if debugModel {
		logLevel = zapcore.DebugLevel
	}
	return zapcore.NewCore(encoder, writerSyncer, logLevel)
}

// newFileCore 文件输出的方式
func newFileCore() zapcore.Core {
	encoder := newFileTextEncoder()

	fileWriteSyncer := newMultiFileWriter()
	writerSyncer := zapcore.NewMultiWriteSyncer(fileWriteSyncer) // 可以指定多个输出

	return zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
}
