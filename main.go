package main

import (
	"awesomeProject/conf"
	"awesomeProject/logger"
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

const version = "1.0.0"

var (
	help           bool
	port           int
	stage          string
	serverAddress  string
	plugins        []server.Plugin
	serviceInfos   []rpc.ServiceInfo
	zipKinReporter reporter.Reporter
)

func init() {
	zipKinReporter = utils.InitTracer(conf.ZipKinHostPort, conf.ZipTag, serverAddress) // 初始化zipKin
	initOptions()
	iniLogger()
	initPlugins()
	initServices()
}

func done() {
	logger.Log.Close()
	zipKinReporter.Close()
}

func initOptions() {
	// 初始化选项
	flag.BoolVar(&help, "h", false, "this help")
	flag.IntVar(&port, "p", 9999, "rpc server port")
	flag.StringVar(&stage, "s", "DEV", "rpc server port")
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr, `
RpcxDemo version: %s
Usage: main [-p port]
Options:
		`, version)
		flag.PrintDefaults()
	}
	serverAddress = fmt.Sprintf("%s:%d", utils.CurrentIp, port)
}

func iniLogger() {
	// 初始化logger
	logger.InitLogger(conf.ProjectName)
}

func initPlugins() {
	// etcd服务注册插件
	etcdRegistryPlugin := p.NewEtcdRegistryPlugin(serverAddress, conf.ServiceBasePath, conf.EtcdUrls)
	plugins = append(plugins, etcdRegistryPlugin)

	// 限流插件
	limiterPlugin := p.NewLimiterPlugin(time.Second, 10000)
	plugins = append(plugins, limiterPlugin)

	// openTracing全链路追踪插件
	openTracingPlugin := p.NewOpenTracingPlugin(conf.ZipKinHostPort, conf.ZipTag, serverAddress)
	plugins = append(plugins, openTracingPlugin)
}

func initServices() {
	// 初始化服务
	serviceInfos = append(serviceInfos, rpc.ServiceInfo{
		ServerName: "Greeter",
		Handler:    new(services.GreeterHandler),
		Metadata:   "group=greeter",
	})
}

func main() {
	defer done()
	flag.Parse()
	if help {
		flag.Usage()
	} else {
		rpcServer := rpc.NewRpcServer(serverAddress) // 初始化RPC服务
		rpcServer.Server.AuthFunc = p.AuthFunc       // 认证插件
		rpcServer.AddPlugins(plugins)                // 添加插件
		rpcServer.RegisterServices(&serviceInfos)    // 注册服务
		rpcServer.Start()                            // 启动RPC服务
	}
}
