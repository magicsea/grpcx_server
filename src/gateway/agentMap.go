package main

import (
	"sync"
)
var agentMap sync.Map // map[userId]*Agent
var agentAddrMap sync.Map //map[addr]userId
func init()  {

}

func addAgent(a *Agent)  {
	agentMap.Store(a.userID,a)
	agentAddrMap.Store(a.clientAddr,a.userID)
}


func removeAgent(a *Agent)  {
	agentMap.Delete(a.userID)
	agentAddrMap.Delete(a.clientAddr)
}


func findAgent(id int64) *Agent  {
	v,ok := agentMap.Load(id)
	if !ok {
		return nil
	}
	return v.(*Agent)
}

func checkAgent(addr string) (int64,bool) {
	v,ok := agentAddrMap.Load(addr)
	if !ok {
		return 0,false
	}
	return v.(int64),true
}

func eachAgentDo(f func(a *Agent))  {
	agentMap.Range(func(key, value interface{}) bool {
		f(value.(*Agent))
		return true
	})
}
//
//func broadcast(msg string)  {
//	agentMap.Range(func(key, value interface{}) bool{
//		value.(*Agent).SendMsg(&pb.StreamResponse{Pt:&pb.StreamPoint{Name:msg}})
//		return true
//	})
//}