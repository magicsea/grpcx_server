package sd

import (
	"errors"
	"fmt"
	"sync"
	"time"

	//"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/mdns"
	"github.com/spf13/viper"
)

var rg registry.Registry
var rs selector.Selector

// RegisterSever
func RegisterSever(conf *viper.Viper, meta map[string]string, serviceNames ...string) error {

	if len(serviceNames) < 1 {
		return errors.New("need serviceName")
	}

	var opts []registry.Option
	opts = append(opts, registry.Timeout(time.Millisecond*100))
	rg = mdns.NewRegistry(opts...)

	//regist
	for _, serviceName := range serviceNames {
		sv := registry.Service{
			Name:     conf.Get(serviceName + ".type").(string),
			Version:  "0.0.1",
			Metadata: meta,
			Nodes: []*registry.Node{
				{
					Id:       serviceName,
					Address:  conf.Get(serviceName + ".listenAddr").(string),
					Metadata: meta,
				},
			},
		}
		err := rg.Register(&sv)
		if err != nil {
			return err
		}
	}

	//watch
	//w, werr := rg.Watch()
	//if werr != nil {
	//	return werr
	//}
	//go registerWatcher(w)
	rs = selector.NewSelector(selector.Registry(rg), selector.SetStrategy(RoundRobin))
	return nil
}

var selectorFuncs sync.Map //map[string]selector.Next{}

//registerWatcher
func registerWatcher(w registry.Watcher) {
	for {
		fmt.Println("#wait watch event...")
		res, nerr := w.Next()
		fmt.Println("#watch event==========>", res, nerr)
		time.Sleep(time.Second * 1)
		selectorFuncs.Delete(res.Service.Name)
	}
}

//获取一个服务 SelectServiceCache
func SelectServiceCache(stype string) (*registry.Node, error) {
	if f, ok := selectorFuncs.Load(stype); ok {
		fun := f.(selector.Next)
		node, err := fun()
		if err != nil {
			return nil, err
		}
		return node, nil
	}

	next, err := rs.Select(stype)
	if err != nil {
		return nil, err
	}

	selectorFuncs.Store(stype, next)
	node, errn := next()
	if errn != nil {
		return nil, errn
	}
	return node, nil
}

//SelectService
func SelectService(stype string) (*registry.Node, error) {
	next, err := rs.Select(stype)
	if err != nil {
		return nil, err
	}

	node, errn := next()
	if errn != nil {
		return nil, errn
	}
	return node, nil
}

//ListService 获取一类服务
func ListService(sv string) []*registry.Node {
	var l []*registry.Node
	svs, err := rg.GetService(sv)
	if err != nil {
		fmt.Println("ListService error:", sv, err)
		return nil
	}
	for _, s := range svs {
		for _, n := range s.Nodes {
			l = append(l, n)
		}
	}

	return l
}

//WatchService 关注一个类型
func WatchService(stype string) {
	rs.Select(stype)
}

//
//func TestOverflow() {
//	//counts := map[string]int{}
//
//	var opts []registry.Option
//	opts = append(opts, registry.Timeout(time.Millisecond*100))
//	rg := etcd.NewRegistry(opts...) //mdns.NewRegistry(opts...)
//
//	sv := registry.Service{Name: "foo",
//		Version: "1.0.1",
//		Nodes: []*registry.Node{
//			{
//				Id:      "test1-1",
//				Address: "0.0.0.0:10001",
//				Metadata: map[string]string{
//					"foo": "bar",
//				},
//			},
//			{
//				Id:      "test1-2",
//				Address: "0.0.0.0:10002",
//				Metadata: map[string]string{
//					"foo": "bar2",
//				},
//			},
//		}}
//
//	rg.Register(&sv)
//
//	rs := selector.NewSelector(selector.Registry(rg), selector.SetStrategy(selector.RoundRobin))
//
//	next, err := rs.Select("foo")
//	if err != nil {
//		fmt.Printf("Unexpected error calling default select: %v\n", err)
//	}
//	next()
//}
