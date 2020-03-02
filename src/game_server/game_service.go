package main

import (
	"context"
	"share"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	. "pb"
)

//GameService
type GameService struct{}

//返回
func (s *GameService) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "GameService.Hello canceled")
	}
	oc, _ := metadata.FromIncomingContext(ctx)
	uid, err := share.GetUIDFromContext(ctx)
	println("Hello.....", req.Request, uid, err, oc)

	return &HelloResponse{Response: req.GetRequest() + " Server"}, err
}

//广播
func (s *GameService) HelloWorld(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "GameService.Hello canceled")
	}
	oc, _ := metadata.FromIncomingContext(ctx)
	uid, err := share.GetUIDFromContext(ctx)
	println("HelloWorld.....", req.Request, uid, err, oc)
	broadcastClients(ctx, "helloworld", req)
	time.Sleep(time.Second)
	return &HelloResponse{Response: req.GetRequest() + " Server"}, err
}

//广播
func (s *GameService) TellYou(ctx context.Context, req *TellRequest) (*TellRsp, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "GameService.Hello canceled")
	}
	oc, _ := metadata.FromIncomingContext(ctx)
	uid, err := share.GetUIDFromContext(ctx)

	println("TellYou.....", req.Request, req.TargetId, uid, err, oc)

	sendClient(ctx, req.TargetId, "tellyou", req)
	time.Sleep(time.Second)
	return &TellRsp{Request: req.Request, TargetId: req.TargetId}, err
}
