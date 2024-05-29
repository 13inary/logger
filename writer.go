package logger

import (
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logFilePath = "./log/service.log"
)

// newConsoleWriter 控制台的输出
func newConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

// newFileWriter 单个文件的输出
func newFileWriter() zapcore.WriteSyncer {
	file, _ := os.Create(logFilePath)
	return zapcore.AddSync(file)
}

// newMultiFileWriter 文件有归档
func newMultiFileWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    2,     // 单个文件最大。单位为M
		MaxBackups: 50,    // 文件最大个数
		MaxAge:     0,     // 文件保留天数。0为无限期
		Compress:   false, // 是否压缩文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
