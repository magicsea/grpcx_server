package share

import (
	"errors"
	"log"
	"time"

	"sd"

	"google.golang.org/grpc"
)

//ServerNode
type ServerNode struct {
	ServerName string
	Conn       interface{}
}

//var peersMap = map[string][]ServerNode{}
//var peerlocker sync.Mutex
//var peersMap = sync.Map{}

//Err_ExistPeer
var Err_ExistPeer = errors.New("Err_ExistPeer")

//func InitPeers(stypes ...string) error {
//
//	for _, stype := range stypes {
//		nodes := GetServersByType(stype)
//		for _, v := range nodes {
//			addr := v.OutAddr
//			conn, err := grpc.Dial(addr, grpc.WithInsecure())
//			if err != nil {
//				log.Fatalf("grpc.Dial err: %v,%v", err, addr)
//				return err
//			}
//			AddPeer(stype, v.Name, conn)
//		}
//	}
//	return nil
//}

//addPeer
//func addPeer(peerType string, peerName string, client interface{}) error {
//
//	if _, ok := peersMap[peerType]; !ok {
//		peersMap[peerType] = []ServerNode{}
//	}
//	for _, node := range peersMap[peerType] {
//		if node.ServerName == peerName {
//			return Err_ExistPeer
//		}
//	}
//
//	peersMap[peerType] = append(peersMap[peerType], ServerNode{peerName, client})
//	return nil
//}

//todo:连接的缓存和健康检查
//GetPeer
func GetOrCreatePeer(id string, addr string) (interface{}, error) {
	//create
	conn, err := grpc.Dial(addr, grpc.WithTimeout(time.Second*2), grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v,%v", err, addr)
		return nil, err
	}
	return conn, nil
}

var callCounter int64

//轮询调用
//func GetPeerRoundRobin(stype string) (interface{}, bool) {
//	nodes := GetPeersByType(stype)
//	if len(nodes) < 1 {
//		return nil, false
//	}
//
//	r := atomic.AddInt64(&callCounter, 1)
//	l := int64(len(nodes))
//	endpoint := nodes[r%l]
//	return endpoint, true
//}

//所有调用一次
func CallEachPeer(stype string, f func(conn interface{})) {
	nodes := ListPeersByType(stype)
	for _, v := range nodes {
		f(v)
	}
}

//SelectPeer 轮询一个节点
func SelectPeer(peerType string) (interface{}, bool) {
	node, err := sd.SelectService(peerType)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	c, erra := GetOrCreatePeer(node.Id, node.Address)
	if erra != nil {
		log.Fatal(erra)
		return nil, false
	}
	return c, true
}

//GetPeersByType 列举某类型所有节点
func ListPeersByType(peerType string) []interface{} {
	nodes := sd.ListService(peerType)
	var list []interface{}
	for _, n := range nodes {
		c, err := GetOrCreatePeer(n.Id, n.Address)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, c)
	}

	return list
}
