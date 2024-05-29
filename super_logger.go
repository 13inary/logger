package logger

func Debugf(template string, args ...interface{}) {
	superLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	superLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	superLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	superLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	superLogger.Fatalf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	superLogger.Panicf(template, args...)
}
