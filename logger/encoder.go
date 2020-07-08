package logger

import (
	"go.uber.org/zap/zapcore"
	"time"
)

var ErrorEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "ts",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "errMsg",
	StacktraceKey:  "errTrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}
var CommonTimeEncoder = func (t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
var CommonEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "ts",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     CommonTimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

var ErrorJsonEncoder = zapcore.NewJSONEncoder(ErrorEncoderConfig)
var ErrorConsoleEncoder = zapcore.NewConsoleEncoder(ErrorEncoderConfig)
var CommonJsonEncoder = zapcore.NewJSONEncoder(CommonEncoderConfig)
var CommonConsoleEncoder = zapcore.NewConsoleEncoder(CommonEncoderConfig)
