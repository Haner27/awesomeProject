package main

import (
	"awesomeProject/conf"
	"awesomeProject/dao/es"
	"awesomeProject/logger"
	"awesomeProject/mq/kafka"
	p "awesomeProject/plugins"
	"awesomeProject/rpc"
	"awesomeProject/services"
	"awesomeProject/utils/iputil"
	"awesomeProject/utils/zipkinutil"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/smallnest/rpcx/server"
	"go.uber.org/zap/zapcore"
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
	zipKinReporter = zipkinutil.InitTracer(conf.ZipKinHostPort, conf.ZipTag, serverAddress) // 初始化zipKin
	initOptions()
	initDao()
	initMq()
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
	serverAddress = fmt.Sprintf("%s:%d", iputil.CurrentIp, port)
}

func iniLogger() {
	// 初始化logger
	core := zapcore.NewTee(
		zapcore.NewCore(logger.CommonConsoleEncoder, logger.StdoutSyncEr, logger.CommonLevelEnable),
		zapcore.NewCore(logger.ErrorJsonEncoder, logger.StdoutSyncEr, logger.ErrorLevelEnable),
		zapcore.NewCore(logger.ErrorJsonEncoder, logger.NewElasticSearchSyncEr(es.EsClient), logger.ErrorLevelEnable),
		zapcore.NewCore(logger.ErrorJsonEncoder, logger.NewKafkaSyncEr(kafka.Producer), logger.ErrorLevelEnable),
	)
	logger.InitLogger(conf.ProjectName, core)
}

func initDao() {
	// 初始化数据访问对象
	es.InitEsDao(conf.EsUrls)  // 初始化es client
}

func initMq() {
	// 生产者配置
	kafkaConfig := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	kafkaConfig.Producer.Partitioner = sarama.NewReferenceHashPartitioner
	// 是否等待成功和失败后的响应
	kafkaConfig.Producer.Return.Successes = true
	// buffer 每隔多少时间触发flush
	kafkaConfig.Producer.Flush.Frequency = 5 * time.Second
	// buffer 最多装多少条消息
	kafkaConfig.Producer.Flush.MaxMessages = 10000
	// buffer 装多少条消息触发flush
	kafkaConfig.Producer.Flush.Messages = 200
	// 初始化消息队列
	kafka.InitProducer(conf.KafkaUrls, conf.KafkaTopic, kafkaConfig)
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
		logger.Log.Errorw("test error log", "name", "hannengfang", "age", 28)
		rpcServer := rpc.NewRpcServer(serverAddress) // 初始化RPC服务
		rpcServer.Server.AuthFunc = p.AuthFunc       // 认证插件
		rpcServer.AddPlugins(plugins)                // 添加插件
		rpcServer.RegisterServices(&serviceInfos)    // 注册服务
		rpcServer.Start()                            // 启动RPC服务
	}
}
