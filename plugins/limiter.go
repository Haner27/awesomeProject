package plugin

import (
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

func NewLimiterPlugin(interval time.Duration, times int64) server.Plugin {
	rateLimiter := serverplugin.NewRateLimitingPlugin(interval, times)
	return rateLimiter
}
