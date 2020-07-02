package rpc

import (
	"awesomeProject/conf"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"time"
)

var xClientOption client.Option

func init() {
	initOption()
}

func initOption() {
	xClientOption = client.DefaultOption
	xClientOption.Heartbeat = true
	xClientOption.HeartbeatInterval = time.Second
	xClientOption.ReadTimeout = 10 * time.Second
	xClientOption.SerializeType = protocol.ProtoBuffer
}

func getEtcdDiscovery(serviceBasePath, servicePath string, etcdAddr []string) *client.ServiceDiscovery {
	d := client.NewEtcdDiscovery(
		serviceBasePath,
		servicePath,
		etcdAddr,
		nil,
	) // Etcd 服务发现
	return &d
}

func NewXClient(serviceBasePath, serviceName string, etcdAddr []string) client.XClient {
	discovery := getEtcdDiscovery(serviceBasePath, serviceName, etcdAddr)
	xClient := client.NewXClient(
		serviceName,
		client.Failtry,
		client.RandomSelect,
		*discovery,
		xClientOption,
	)
	xClient.Auth(conf.AuthToken)
	xClient.GetPlugins().Add(&client.OpenTracingPlugin{})// 认证token
	return xClient
}
