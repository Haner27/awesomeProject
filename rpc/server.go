package rpc

import (
	"github.com/smallnest/rpcx/server"
	"log"
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
	log.Print("[services]:\n")
	for _, serviceInfo := range *serviceInfos {
		err := rs.Server.RegisterName(serviceInfo.ServerName, serviceInfo.Handler, serviceInfo.Metadata)
		if err != nil {
			log.Fatalf("Register %s service failed.", serviceInfo.ServerName)
		} else {
			log.Printf("Register %s service.\n", serviceInfo.ServerName)
		}
	}
}

func (rs *Server) AddPlugins(plugins []server.Plugin) {
	// 添加插件
	log.Print("[plugins]:\n")
	for _, plugin := range plugins {
		t := reflect.TypeOf(plugin)
		rs.Server.Plugins.Add(plugin)
		log.Printf("Add %s plugin.", t.Elem().Name())
	}
}

func (rs *Server) Start() {
	// 启动服务
	log.Println("[starting]:")
	log.Println(rs.Address)
	err := rs.Server.Serve("tcp", rs.Address)
	if err != nil {
		log.Fatal("Server start failed.")
	}
}
