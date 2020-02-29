package main

import (
	"context"
	"errors"

	"log"

	. "pb"
)

//=======================================================================
type GateService struct{}

func (s *GateService) KickAgent(ctx context.Context, req *KickAgentReq) (*Rsp, error) {
	log.Printf("KickAgent:%+v", req)
	return &Rsp{Result: 0}, nil
}

func (s *GateService) PushClient(ctx context.Context, req *PushClientReq) (*Rsp, error) {
	log.Printf("PushClient:%+v", req)
	a := findAgent(req.Uid)
	if a == nil {
		return &Rsp{10001}, errors.New("not found user")
	}
	a.SendRaw(req.Data)
	return &Rsp{Result: 0}, nil
}

func (s *GateService) BroadcastClient(ctx context.Context, req *BroadcastClientReq) (*Rsp, error) {
	log.Printf("BroadcastClient:%+v", req)
	eachAgentDo(func(a *Agent) {
		a.SendRaw(req.Data)
	})
	return &Rsp{Result: 0}, nil
}
