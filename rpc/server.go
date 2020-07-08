package rpc

import (
	"awesomeProject/logger"
	"github.com/smallnest/rpcx/server"
	"reflect"
)

type ServiceInfo struct {
	ServerName string
	Handler    interface{}
	Metadata   string
}

type Server struct {
	Address string
	Server  *server.Server
}

func NewRpcServer(address string) *Server {
	rpcServer := &Server{
		Address: address,
		Server:  server.NewServer(),
	}
	return rpcServer
}

func (rs *Server) RegisterServices(serviceInfos *[]ServiceInfo) {
	logger.Log.Info("【SERVICES】")
	for _, serviceInfo := range *serviceInfos {
		err := rs.Server.RegisterName(serviceInfo.ServerName, serviceInfo.Handler, serviceInfo.Metadata)
		if err != nil {
			logger.Log.Errorw("Registered a service failed.", "serviceName", serviceInfo.ServerName)
		} else {
			logger.Log.Infow("Registered a service.", "serviceName", serviceInfo.ServerName)
		}
	}
}

func (rs *Server) AddPlugins(plugins []server.Plugin) {
	// 添加插件
	logger.Log.Info("【PLUGINS】")
	for _, plugin := range plugins {
		t := reflect.TypeOf(plugin)
		rs.Server.Plugins.Add(plugin)
		logger.Log.Infow("Added a plugin.", "pluginName", t.Elem().Name())
	}
}

func (rs *Server) Start() {
	// 启动服务
	logger.Log.Infow("【STARTING SERVER】", "serverAddress", rs.Address)
	err := rs.Server.Serve("tcp", rs.Address)
	if err != nil {
		logger.Log.Error("Server start failed.")
	}
}
