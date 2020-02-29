package main

import (
	"errors"
	"sync"
)

var peersMap = map[string]map[string]interface{}{}
var peerlocker sync.Mutex

var Err_ExistPeer  = errors.New("Err_ExistPeer")

//AddPeer
func AddPeer(peerType string,peerName string,client interface{}) error {
	//conn, err := grpc.Dial(addr)
	//if err != nil {
	//	log.Fatalf("grpc.Dial err: %v,%v,%v", err,peerType,addr)
	//	return err
	//}

	peerlocker.Lock()
	defer  peerlocker.Unlock()

	if _,ok:=peersMap[peerType];!ok {
		peersMap[peerType] = make(map[string]interface{})
	}

	if _,ok := peersMap[peerType][peerName];ok {
		return Err_ExistPeer
	}
	peersMap[peerType][peerName] = client
	return nil
}

//GetPeersByType
func GetPeersByType(peerType string) []interface{} {
	peerlocker.Lock()
	defer  peerlocker.Unlock()

	m,ok := peersMap[peerType]
	if !ok {
		return nil
	}

	var list []interface{}
	for _, v := range m {
		list = append(list,v)
	}

	return list
}

//GetPeer
func GetPeer(peerType string,peerName string) (interface{},bool) {
	peerlocker.Lock()
	defer  peerlocker.Unlock()

	m,ok := peersMap[peerType]
	if !ok {
		return nil,false
	}
	v,found := m[peerName]
	return v,found
}