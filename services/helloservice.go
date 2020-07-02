package services

import (
	"awesomeProject/conf"
	"awesomeProject/proto/calculate"
	helloWorld "awesomeProject/proto/helloword"
	"awesomeProject/rpc"
	"context"
	"fmt"
)

type GreeterHandler struct{}

// SayHello is server rpc method as defined
func (s *GreeterHandler) SayHello(ctx context.Context, args *helloWorld.HelloRequest, reply *helloWorld.HelloReply) (err error) {
	xClient := rpc.NewXClient(conf.ServiceBasePath2, "Calculate", conf.EtcdUrls)
	defer xClient.Close()
	client := calculate.NewCalculateClient(xClient)
	calRequest := &calculate.CalRequest{
		X: 100,
		Y: 50,
	}
	calReply1, err := client.Add(ctx, calRequest)
	if err != nil {
		return fmt.Errorf("SayHello call add failed:%v", err)
	}
	calReply2, err := client.Sub(ctx, calRequest)
	if err != nil {
		return fmt.Errorf("SayHello call sub failed:%v", err)
	}
	calReply3, err := client.Mul(ctx, calRequest)
	if err != nil {
		return fmt.Errorf("SayHello call mul failed:%v", err)
	}
	calReply4, err := client.Div(ctx, calRequest)
	if err != nil {
		return fmt.Errorf("SayHello call div failed:%v", err)
	}

	reply.Message = fmt.Sprintf("hello %s! \n" +
		"%[2]d + %[3]d = %[4]d\n" +
		"%[2]d + %[3]d = %[5]d\n" +
		"%[2]d + %[3]d = %[6]d\n" +
		"%[2]d + %[3]d = %[7]d\n" +
		"", args.Name, calRequest.X, calRequest.Y, calReply1.R, calReply2.R, calReply3.R, calReply4.R)
	return nil
}
