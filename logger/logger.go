package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *Logger
var zapLogger *zap.Logger

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(name string) *Logger {
	core := zapcore.NewTee(
		zapcore.NewCore(CommonConsoleEncoder, StdoutSyncEr, CommonLevelEnable),
		zapcore.NewCore(ErrorJsonEncoder, StdoutSyncEr, ErrorLevelEnable),
	)
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

func InitLogger(name string) {
	Log = NewLogger(name)
}