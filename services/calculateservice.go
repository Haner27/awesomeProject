package services

import (
	"awesomeProject/proto/calculate"
	"context"
)

type CalculateHandler struct {}

func (s *CalculateHandler) Add(ctx context.Context, args *calculate.CalRequest, reply *calculate.CalReply) (err error) {
	reply.R = args.X + args.Y
	return nil
}

// Sub is server rpc method as defined
func (s *CalculateHandler) Sub(ctx context.Context, args *calculate.CalRequest, reply *calculate.CalReply) (err error) {
	reply.R = args.X - args.Y
	return nil
}

// Mul is server rpc method as defined
func (s *CalculateHandler) Mul(ctx context.Context, args *calculate.CalRequest, reply *calculate.CalReply) (err error) {
	reply.R = args.X * args.Y
	return nil
}

// Div is server rpc method as defined
func (s *CalculateHandler) Div(ctx context.Context, args *calculate.CalRequest, reply *calculate.CalReply) (err error) {
	reply.R = args.X / args.Y
	return nil
}


