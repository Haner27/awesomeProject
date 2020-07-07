package plugin

import (
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

func NewEtcdRegistryPlugin(serviceAddress, serviceBasePath string, etcdUrls []string) server.Plugin {
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: serviceAddress,
		EtcdServers:    etcdUrls,
		BasePath:       serviceBasePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	return r
}
