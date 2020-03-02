package share

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc"
)

type ServerNode struct {
	ServerName string
	Conn       interface{}
}

var peersMap = map[string][]ServerNode{}
var peerlocker sync.Mutex

var Err_ExistPeer = errors.New("Err_ExistPeer")

func InitPeers(stypes ...string) error {
	for _, stype := range stypes {
		nodes := GetServersByType(stype)
		for _, v := range nodes {
			addr := v.OutAddr
			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("grpc.Dial err: %v,%v", err, addr)
				return err
			}
			AddPeer(stype, v.Name, conn)
		}
	}
	return nil
}

//AddPeer
func AddPeer(peerType string, peerName string, client interface{}) error {
	//conn, err := grpc.Dial(addr)
	//if err != nil {
	//	log.Fatalf("grpc.Dial err: %v,%v,%v", err,peerType,addr)
	//	return err
	//}

	peerlocker.Lock()
	defer peerlocker.Unlock()

	if _, ok := peersMap[peerType]; !ok {
		peersMap[peerType] = []ServerNode{}
	}
	for _, node := range peersMap[peerType] {
		if node.ServerName == peerName {
			return Err_ExistPeer
		}
	}

	peersMap[peerType] = append(peersMap[peerType], ServerNode{peerName, client})
	return nil
}

//GetPeersByType
func GetPeersByType(peerType string) []interface{} {
	peerlocker.Lock()
	defer peerlocker.Unlock()

	m, ok := peersMap[peerType]
	if !ok {
		return nil
	}

	var list []interface{}
	for _, v := range m {
		list = append(list, v.Conn)
	}

	return list
}

//GetPeer
func GetPeer(peerType string, peerName string) (interface{}, bool) {
	peerlocker.Lock()
	defer peerlocker.Unlock()

	m, ok := peersMap[peerType]
	if !ok {
		return nil, false
	}
	for _, node := range m {
		if node.ServerName == peerName {
			return node.Conn, true
		}
	}

	return nil, false
}

var callCounter int64

//轮询调用
func GetPeerRoundRobin(stype string) (interface{}, bool) {
	nodes := GetPeersByType(stype)
	if len(nodes) < 1 {
		return nil, false
	}

	r := atomic.AddInt64(&callCounter, 1)
	l := int64(len(nodes))
	endpoint := nodes[r%l]
	return endpoint, true
}

//所有调用一次
func CallEachPeer(stype string, f func(conn interface{})) {
	nodes := GetPeersByType(stype)
	for _, v := range nodes {
		f(v)
	}
}
