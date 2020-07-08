package logger

import (
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
)

var DevNullSyncEr = zapcore.AddSync(ioutil.Discard)
var StdoutSyncEr = zapcore.Lock(os.Stdout)
var StderrSyncEr = zapcore.Lock(os.Stderr)
