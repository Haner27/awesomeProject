package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ErrorLevelEnable = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
})
var CommonLevelEnable = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
})
