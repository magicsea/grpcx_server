package main

import (
	"context"
	"flag"
	"pb"
	"share"
	"github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"log"
	"net"
)




var listen = flag.String("listen",":9001","listen addr")
func main() {
	flag.Parse()

	initClient()

	server := grpc.NewServer()
	pb.RegisterGameServiceServer(server, &GameService{})


	println("listen  ",*listen)
	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

func initClient() error {
	addr := "127.0.0.1:3002"
	conn, err := grpc.Dial(addr,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v,%v", err,addr)
		return err
	}
	gateway := pb.NewGateServiceClient(conn)
	AddPeer(share.GateWay,"gate1",gateway)
	return nil
}

func broadcastClients(ctx context.Context,msgId string,msg proto.Message) error {
	gates := GetPeersByType(share.GateWay)
	data,err:= proto.Marshal(msg)
	if err != nil {
	    return err
	}
	raw := pb.RawMsg{MsgId:msgId,MsgData:data}
	_ = raw
	req := pb.BroadcastClientReq{Data:&raw}
	for _, v := range gates {
		client := v.(pb.GateServiceClient)
		rsp,callErr := client.BroadcastClient(ctx,&req)
		_ = rsp
		if callErr != nil {
		    log.Println("send gateway error:",callErr)
		}
	}

	return nil
}

//todo:这里应该从cache里找到目标服务器，测试临时广播下
func sendClient(ctx context.Context,uid int64,msgId string,msg proto.Message) error {
	gates := GetPeersByType(share.GateWay)
	data,err:= proto.Marshal(msg)
	if err != nil {
		return err
	}
	raw := pb.RawMsg{MsgId:msgId,MsgData:data}
	_ = raw
	req := pb.PushClientReq{Uid:uid,Data:&raw}
	for _, v := range gates {
		client := v.(pb.GateServiceClient)
		rsp,callErr := client.PushClient(ctx,&req)
		_ = rsp
		if callErr != nil {
			log.Println("send gateway error:",callErr)
		}
	}

	return nil
}


//
//func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	newc := context.WithValue(ctx,"kk",11)
//	log.Printf("gRPC method: %s, %v", info.FullMethod, req)
//	resp, err := handler(newc, req)
//	log.Printf("gRPC method: %s, %v", info.FullMethod, resp)
//	return resp, err
//}
//
//func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
//	defer func() {
//		if e := recover(); e != nil {
//			debug.PrintStack()
//			err = status.Errorf(codes.Internal, "Panic err: %v", e)
//		}
//	}()
//
//	return handler(ctx, req)
//}
//opts := []grpc.ServerOption{
////grpc.Creds(c),
////grpc.CustomCodec(RawCodec()),
//grpc.UnknownServiceHandler(func(srv interface{}, stream grpc.ServerStream) error {
//// little bit of gRPC internals never hurt anyone
//
//fullMethodName, ok := grpc.MethodFromServerStream(stream)
//if !ok {
//return grpc.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
//}
//md := stream.Context()
//_ = md
//f := &pb.SearchRequest{}
//if err := stream.RecvMsg(f); err != nil {
//return err
//}
//err2 := stream.SendMsg(&pb.SearchResponse{Response:f.Request})
//
//fmt.Println("recv:",fullMethodName,f,err2)
//return nil
//} ),
//grpc_middleware.WithUnaryServerChain(
//RecoveryInterceptor,
//LoggingInterceptor,
//),
//}