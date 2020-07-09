package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *Logger
var zapLogger *zap.Logger

func init() {
	l, _ := zap.NewDevelopment()
	Log = &Logger{
		l.Sugar(),
	}
}

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(name string, core zapcore.Core) *Logger {
	zapLogger = zap.New(
		core,
		//zap.AddCaller(),
		zap.AddStacktrace(ErrorLevelEnable),
		//zap.Hooks(FilterStage),
	)
	sugar := zapLogger.Sugar()
	sugar.Named(name)
	return &Logger{
		sugar,
	}
}

func (l *Logger) Close() {
	zapLogger.Sync()
}

func InitLogger(name string, core zapcore.Core) {
	Log = NewLogger(name, core)
}