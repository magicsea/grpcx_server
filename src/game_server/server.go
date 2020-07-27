package main

import (
	"context"
	"log"
	"net"
	"pb"
	"share"

	"net/http"
	_ "net/http/pprof"

	"github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
	share.Init("game1")
	//sd.WatchService(share.GateWay)
	//sd.TestOverflow()
	server := grpc.NewServer()
	pb.RegisterGameServiceServer(server, &GameService{})

	addr := share.GetServerConf("listenAddr").(string)
	println("listen  ", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	go http.ListenAndServe(":7070", nil)
	server.Serve(lis)
}

func broadcastClients(ctx context.Context, msgId string, msg proto.Message) error {

	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	raw := pb.RawMsg{MsgId: msgId, MsgData: data}
	req := pb.BroadcastClientReq{Data: &raw}

	share.CallEachPeer(share.GateWay, func(conn interface{}) {
		client := pb.NewGateServiceClient(conn.(*grpc.ClientConn))
		rsp, callErr := client.BroadcastClient(ctx, &req)
		_ = rsp
		if callErr != nil {
			log.Println("send gateway error:", callErr)
		}
	})

	return nil
}

//todo:这里应该从cache里找到目标服务器，测试临时广播下
func sendClient(ctx context.Context, uid int64, msgId string, msg proto.Message) error {

	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	raw := pb.RawMsg{MsgId: msgId, MsgData: data}
	req := pb.PushClientReq{Uid: uid, Data: &raw}

	share.CallEachPeer(share.GateWay, func(conn interface{}) {
		client := pb.NewGateServiceClient(conn.(*grpc.ClientConn))
		rsp, callErr := client.PushClient(ctx, &req)
		_ = rsp
		if callErr != nil {
			log.Println("send gateway error:", callErr)
		}
	})
	return nil
}
