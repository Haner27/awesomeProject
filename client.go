package main

import (
	"awesomeProject/conf"
	helloWorld "awesomeProject/proto/helloword"
	"awesomeProject/rpc"
	"awesomeProject/utils"
	"context"
	"fmt"
)

func main() {
	init := utils.InitTracer(conf.ZipKinHostPort, "third-client", utils.CurrentIp) // zipkin
	defer init.Close()

	xClient := rpc.NewXClient(conf.ServiceBasePath, "Greeter", conf.EtcdUrls)
	defer xClient.Close()

	client := helloWorld.NewGreeterClient(xClient)
	for i:=0;i<10;i++ {
		args := &helloWorld.HelloRequest{
			Name: "rpcx",
		}
		reply, err := client.SayHello(context.Background(), args)
		if err != nil {
			panic(err)
		}
		fmt.Println("reply: ", reply.Message)
	}
}
