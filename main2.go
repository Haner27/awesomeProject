package main

import (
	"awesomeProject/conf"
	p "awesomeProject/plugins"
	"awesomeProject/rpc"
	"awesomeProject/services"
	"awesomeProject/utils"
	"flag"
	"fmt"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/smallnest/rpcx/server"
	"os"
	"time"
)

const version2 = "1.0.0"

var (
	help2           bool
	port2           int
	serverAddress2  string
	plugins2        []server.Plugin
	serviceInfos2   []rpc.ServiceInfo
	zipKinReporter2 reporter.Reporter
)

func init() {
	zipKinReporter2 = utils.InitTracer(conf.ZipKinHostPort2, conf.ZipTag2, serverAddress2) // 初始化zipKin
	initOptions2()
	initPlugins2()
	initServices2()
}

func done2() {
	zipKinReporter2.Close()
}

func initOptions2() {
	// 初始化选项
	flag.BoolVar(&help2, "h", false, "this help")
	flag.IntVar(&port2, "p", 8888, "rpc server port")
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr, `
RpcxDemo version: %s
Usage: main [-p port]
Options:
		`, version2)
		flag.PrintDefaults()
	}
	serverAddress2 = fmt.Sprintf("%s:%d", utils.CurrentIp, port2)
}

func initPlugins2() {
	// etcd服务注册插件
	etcdRegistryPlugin2 := p.NewEtcdRegistryPlugin(serverAddress2, conf.ServiceBasePath2, conf.EtcdUrls2)
	plugins2 = append(plugins2, etcdRegistryPlugin2)

	// 限流插件
	limiterPlugin2 := p.NewLimiterPlugin(time.Second, 10000)
	plugins2 = append(plugins2, limiterPlugin2)

	// openTracing全链路追踪插件
	openTracingPlugin2 := p.NewOpenTracingPlugin(conf.ZipKinHostPort2, conf.ZipTag2, serverAddress2)
	plugins2 = append(plugins2, openTracingPlugin2)
}

func initServices2() {
	// 初始化服务
	serviceInfos2 = append(serviceInfos2, rpc.ServiceInfo{
		ServerName: "Calculate",
		Handler:    new(services.CalculateHandler),
		Metadata:   "group=calculator",
	})
}

func main() {
	defer done2()
	flag.Parse()
	if help2 {
		flag.Usage()
	} else {
		rpcServer2 := rpc.NewRpcServer(serverAddress2) // 初始化RPC服务
		rpcServer2.Server.AuthFunc = p.AuthFunc        // 认证插件
		rpcServer2.AddPlugins(plugins2)                // 添加插件
		rpcServer2.RegisterServices(&serviceInfos2)    // 注册服务
		rpcServer2.Start()                             // 启动RPC服务
	}
}
