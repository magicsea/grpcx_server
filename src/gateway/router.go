package main

import (
	"errors"
	"fmt"

	"time"

	"share"

	"sd"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	MODULE_TEXT string = "module"
	UserID_Key  string = "uid"
)

func inSlice(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}

//RoundRobinConfigRouter
func RoundRobinConfigRouter(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
	//验证链接
	p, _ := peer.FromContext(ctx)
	userId, found := checkAgent(p.Addr.String())
	if !found {
		fmt.Println("user session not exist:", p.Addr)
		return nil, nil, grpc.Errorf(codes.Internal, "user session not exist")
	}
	//md,_ := metadata.FromIncomingContext(serverStream.Context())
	fmt.Printf("RoundRobinConfigRouter...................addr=%v  uid=%v\n", p, userId)

	// 获取配置信息
	var rule, foundRule = share.GetRuleMatchProto(fullMethodName)
	if !foundRule {
		println("not found rule:", fullMethodName)
		return nil, nil, errors.New("not found rule")
	}

	// 仅转发外部的请求
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := fmt.Errorf("incoming metadata is empty\n")
		println(err.Error())
		return nil, nil, err
	}
	md[UserID_Key] = []string{fmt.Sprintf("%v", userId)}
	outCtx, _ := context.WithCancel(ctx)
	outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())

	retryNum := 3
	// 重试处理
	for i := 0; i <= retryNum; i++ {
		conn, err := getBackendByRR(ctx, rule.Type)
		if err == nil {
			return outCtx, conn, nil
		}
		println("Error=>getBackendByRR failed:", err.Error())
	} //for retryNum

	return nil, nil, errors.New(fmt.Sprintf("pick backend fail:%+v", rule.Type))
}

var rrCounter uint64

//均衡负载
func getBackendByRR(ctx context.Context, serverType string) (*grpc.ClientConn, error) {
	//pick
	node, err := sd.SelectService(serverType)
	if err != nil {
		return nil, err
	}
	//r := atomic.AddUint64(&rrCounter, 1)
	//l := uint64(len(endpoints))
	//endpoint := endpoints[r%l]
	fmt.Printf("balanced, redirecting to [%+v]\n", node)
	//connect
	// 根据获取到的 endpoint, 建立到目的方的 connection
	// 同时, 需要配置客户端 codec 为我们自定义的 codec
	conn, err := grpc.DialContext(ctx, node.Address, grpc.WithTimeout(time.Second*2), grpc.WithBlock(), grpc.WithCodec(RawCodec()), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc.DialContext failed:%v\n", err.Error())
		return nil, err
	}
	return conn, nil
}
