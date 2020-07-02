package plugin

import (
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

func NewOpenTracingPlugin(zipKinHostPort, serverName, hostPort string) server.Plugin {
	zipKinPlugin := &serverplugin.OpenTracingPlugin{}
	return zipKinPlugin
}
