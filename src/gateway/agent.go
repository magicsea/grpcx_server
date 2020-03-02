package main

import (
	"context"
	"fmt"
	"log"
	"pb"
	"sync"
	"sync/atomic"
	"time"

	"share"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/peer"
)

// Agnet
type Agent struct {
	uuid       uuid.UUID
	userID     int64
	clientAddr string
	closeSign  sync.WaitGroup
	closeFunc  context.CancelFunc
	pushStream pb.LoginService_LoginServer
	timer      int64
	hbTicker   int64
}

// NewAgent
func NewAgent(id int64, stream pb.LoginService_LoginServer) *Agent {
	uid, _ := uuid.NewUUID()
	p, _ := peer.FromContext(stream.Context())

	a := Agent{pushStream: stream, userID: id, uuid: uid, clientAddr: p.Addr.String()}
	a.init()
	return &a
}

func (agent *Agent) init() {
	agent.closeSign = sync.WaitGroup{}
	agent.closeSign.Add(1)
	agent.KeepHB()
}

//心跳
func (agent *Agent) KeepHB() {
	atomic.StoreInt64(&agent.hbTicker, time.Now().Unix())
}

func (agent *Agent) run() {
	ticker := time.NewTicker(time.Second * 5)

	ctx, cancelCtx := context.WithCancel(context.Background())
	agent.closeFunc = cancelCtx
loop:
	for {
		select {
		case <-ticker.C:
			agent.onTick()
		case <-ctx.Done():
			println("agent done:", agent.String())
			break loop
		}
	}
	agent.onClose()
	agent.closeSign.Done()
}

func (agent *Agent) onTick() {
	hb := atomic.LoadInt64(&agent.hbTicker)
	t := share.GetServerConf("HeartbeatTime")
	if time.Now().Unix()-hb > t.(int64) {
		agent.Close("heartbeat too long")
		return
	}

	agent.timer++
	println("ticker:", agent.userID, agent.timer, " hb:", hb)

	//err := agent.SendMsg("server_heartbeat",&pb.HeartBeatMsg{Ticker:time.Now().Unix()})
	//if err!=nil {
	//	agent.Close("send heartbeat miss")
	//}
}

func (agent *Agent) onClose() {
	println("agent onClose:", agent.String())
	removeAgent(agent)
}

// Close
func (agent *Agent) Close(reason string) {
	fmt.Println("Close agent:", agent.String(), reason)
	agent.closeFunc()
}

// WaitClose
func (agent *Agent) WaitClose() {
	agent.closeSign.Wait()
}

func (agent *Agent) String() string {
	return fmt.Sprintf("Agent:%v,%v", agent.uuid, agent.userID)
}

//goroutine safe
func (agent *Agent) SendMsg(msgId string, msgData proto.Message) error {
	fmt.Printf("sendmsg:%+v\n", msgData)
	bt, errM1 := proto.Marshal(msgData)
	if errM1 != nil {
		return errM1
	}
	raw := pb.RawMsg{MsgId: msgId, MsgData: bt}
	return agent.SendRaw(&raw)
}

//goroutine safe
func (agent *Agent) SendRaw(raw *pb.RawMsg) error {
	if err := agent.pushStream.Send(&pb.PushMsg{Msg: raw}); err != nil {
		log.Printf("ERROR: sendmsg %+v,%v", raw, err)
		return err
	}
	return nil
}
