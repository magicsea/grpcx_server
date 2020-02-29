package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/gogo/protobuf/proto"

	"google.golang.org/grpc"

	"pb"
)

const (
	//gateway地址
	Addr = ":3001"
)

var clientconn *grpc.ClientConn

func main() {
	println("conn:",Addr)
	conn, err := grpc.Dial(Addr, grpc.WithTimeout(time.Second*3), grpc.WithBlock(), grpc.WithInsecure())
	clientconn = conn
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	println("new client")
	loginClient := pb.NewLoginServiceClient(conn)
	gameClint := pb.NewGameServiceClient(conn)

	println("conn ok,wait")
	println("wait end")
	stream, err := loginClient.Login(context.Background(), &pb.LoginRequest{Token: "aaa", Uid: 1})
	if err != nil {
		log.Fatalf("Login err: %v", err)
	}

	println("login ok!")

	var ch = make(chan pb.RawMsg,128)
	go startCmd(gameClint)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("stream  break!")
				break
			}
			if err != nil {
				log.Fatal("recv error:", err)
				return
			}

			log.Printf("resp: %+v", resp.Msg)
			ch<-*resp.Msg
		}
	}()

	cancel,cancelFunc:= context.WithCancel(context.Background())
	//main loop
	ticker := time.NewTicker(time.Second)
	for ; ;  {
		select {
		case <-ticker.C:
			hb++
			//fmt.Println("ticker ",hb)

			if isLogin {
				//心跳
				tmCtx,_ := context.WithTimeout(context.Background(),time.Second*10)
				rsp,err := loginClient.HeartBeat(tmCtx,&pb.HeartBeatMsg{hb})
				_ = rsp
				if err!=nil {
					log.Fatal("HeartBeat error:",err)
					cancelFunc()
					return
				}
				//fmt.Println("hb rsp:",rsp)
			}
		case msg:=<-ch:
			onRecvMsg(&msg)
		case <-cancel.Done():
			println("GameOver")
			return
		}
	}

}
var hb int64
var isLogin = false
func onRecvMsg(msg *pb.RawMsg) {
	switch msg.MsgId {
	case "loginRsp":
		var rsp = pb.LoginRsp{}
		err := proto.Unmarshal(msg.MsgData, &rsp)
		fmt.Println("push=>loginRsp:", rsp, err)
		isLogin = true
	case "helloworld":
		var rsp = pb.HelloRequest{}
		err := proto.Unmarshal(msg.MsgData, &rsp)
		fmt.Println("push=>helloworld:", rsp, err)
	case "tellyou":
		var rsp = pb.TellRequest{}
		err := proto.Unmarshal(msg.MsgData, &rsp)
		fmt.Println("push=>tellyou:", rsp.Request, err)

	}
}

func startCmd(gameClient pb.GameServiceClient) {
	var cmd, args string
	for {
		fmt.Println("input cmd:")
		fmt.Scanf("%s %s", &cmd, &args)
		tmCtx,_ := context.WithTimeout(context.Background(),time.Second*10)
		if cmd == "hello" {
			rsp, err := gameClient.Hello(tmCtx, &pb.HelloRequest{Request: args})
			fmt.Println("Hello rsp:", rsp.Response, err)
		} else if cmd == "hi" {
			rsp, err := gameClient.HelloWorld(tmCtx, &pb.HelloRequest{Request: args})
			fmt.Println("HelloWorld rsp:", rsp.Response, err)
		} else if cmd == "tellyou" {
			tid,_:= strconv.Atoi(args)
			rsp, err := gameClient.TellYou(tmCtx, &pb.TellRequest{Request: "xyz",TargetId:int64(tid)})
			fmt.Println("tell rsp:", rsp.TargetId, err)
		}


	}

}
