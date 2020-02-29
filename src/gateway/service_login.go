package main

import (
	"context"
	"fmt"
	. "pb"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var loginLocker sync.Mutex

func init() {
	loginLocker = sync.Mutex{}
}

//=======================================================================
type Loginservice struct{}

//登录
func (s *Loginservice) Login(r *LoginRequest, stream LoginService_LoginServer) error {

	st := grpc.ServerTransportStreamFromContext(stream.Context())
	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "not found peer's address")
	}
	fmt.Printf("Login...................%p  st=%p %p %v\n", stream.(grpc.ServerStream), st, p, p.Addr)

	loginLocker.Lock()
	//kick same user
	id := int64(r.Uid)
	a := findAgent(id)
	if a != nil {
		a.Close("kick same user:" + a.String())
		a.WaitClose()
		println("kick ok:" + a.String())
	}

	println("new agent start:", id)
	//new
	newAgent := NewAgent(id, stream)
	addAgent(newAgent)
	//登录成功
	newAgent.SendMsg("loginRsp", &LoginRsp{Name: newAgent.uuid.String(), Uid: newAgent.userID, Result: 0})
	//broadcast("user login=>"+newAgent.String())
	go newAgent.run()

	loginLocker.Unlock()

	newAgent.WaitClose()
	println("agent end:", newAgent.String())
	return nil
}

//心跳
func (s *Loginservice) HeartBeat(ctx context.Context, hb *HeartBeatMsg) (*HeartBeatMsg, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "not found peer's address")
	}
	userId, found := checkAgent(p.Addr.String())
	if !found {
		fmt.Println("user session not exist:", p.Addr)
		return nil, status.Errorf(codes.Internal, "user id not exist")
	}
	user := findAgent(userId)
	if user == nil {
		return nil, status.Errorf(codes.Internal, "user session not exist")
	}
	user.KeepHB()
	return hb, nil
}
