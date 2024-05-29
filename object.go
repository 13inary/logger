package logger

import "go.uber.org/zap/zapcore"

type ObjectBase struct {
}

// 实现 zapcore.ObjectMarshaler 接口
func (u ObjectBase) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("msg", "need implement")
	return nil
}
