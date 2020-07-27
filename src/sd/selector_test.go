package sd

import (
	"fmt"
	"testing"
	"time"

	"github.com/micro/go-micro/v2/registry/mdns"
	"github.com/spf13/viper"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
)

func TestMdns(t *testing.T) {
	//counts := map[string]int{}

	var opts []registry.Option
	opts = append(opts, registry.Timeout(time.Millisecond*100))
	rg := mdns.NewRegistry(opts...)
	sv := registry.Service{
		Name:    "test1",
		Version: "1.0.1",
		Nodes: []*registry.Node{
			{
				Id:      "test1-1",
				Address: "0.0.0.0:10001",
				Metadata: map[string]string{
					"foo": "bar",
				},
			},
		},
	}

	errReg := rg.Register(&sv)
	if errReg != nil {
		t.Fatal(errReg)
	}

	time.Sleep(time.Millisecond * 100)
	svs, err := rg.GetService("test1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("getservice count:", len(svs))
	for k, v := range svs {
		t.Logf("getservice:%v=> %+v", k, v)
	}

}

func TestMdnsSelector(t *testing.T) {
	counts := map[string]int{}

	var opts []registry.Option
	opts = append(opts, registry.Timeout(time.Millisecond*100))
	rg := mdns.NewRegistry(opts...)

	sv := registry.Service{Name: "foo",
		Version: "1.0.1",
		Nodes: []*registry.Node{
			{
				Id:      "test1-1",
				Address: "0.0.0.0:10001",
				Metadata: map[string]string{
					"foo": "bar",
				},
			},
			{
				Id:      "test1-2",
				Address: "0.0.0.0:10002",
				Metadata: map[string]string{
					"foo": "bar2",
				},
			},
		}}

	sv2 := registry.Service{Name: "foo",
		Version: "1.0.1",
		Nodes: []*registry.Node{
			{
				Id:      "test2-1",
				Address: "0.0.0.0:20001",
				Metadata: map[string]string{
					"foo": "bar3",
				},
			},
			{
				Id:      "test2-2",
				Address: "0.0.0.0:20002",
				Metadata: map[string]string{
					"foo": "bar4",
				},
			},
		}}

	rg.Register(&sv)

	w, werr := rg.Watch()
	if werr != nil {
		t.Fatal(werr)
	}
	defer w.Stop()

	//go func() {
	//	res, nerr := w.Next()
	//	fmt.Println("watch:", res, nerr)
	//}()
	time.Sleep(time.Millisecond * 100)

	rs := selector.NewSelector(selector.Registry(rg), selector.SetStrategy(selector.RoundRobin))

	next, err := rs.Select("foo")
	if err != nil {
		t.Errorf("Unexpected error calling default select: %v", err)
	}

	for i := 0; i < 20; i++ {
		if i == 10 {
			rg.Register(&sv2)
			//res, nerr := w.Next()
			//t.Log("watch:", res, nerr)
			time.Sleep(time.Millisecond * 1000)
			//rs2 := selector.NewSelector(selector.Registry(rg), selector.SetStrategy(selector.RoundRobin))

			next, err = rs.Select("foo")
			if err != nil {
				t.Errorf("Unexpected error calling default select: %v", err)
			}
			//t.Log("reg2")
		}
		//time.Sleep(time.Millisecond * 100)
		node, err := next()
		if err != nil {
			t.Errorf("Expected node err, got err: %v", err)
		}
		counts[node.Id]++
	}

	t.Logf("Default Counts %v", counts)
	time.Sleep(time.Second * 1)
}

func TestRegistServer(t *testing.T) {
	conf := viper.New()
	conf.Set("game1.listenAddr", "0.0.0.0:3001")
	conf.Set("game1.type", "game")
	conf.Set("game2.listenAddr", "0.0.0.0:3002")
	conf.Set("game2.type", "game")
	err := RegisterSever(conf, nil, "game1", "game2")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		node, errs := SelectService("game")
		if errs != nil {
			t.Fatal(errs)
		}
		t.Log("found:", node)
	}

}

func TestRegistGameOne(t *testing.T) {
	conf := viper.New()
	conf.Set("gate1.listenAddr", "0.0.0.0:3001")
	conf.Set("gate1.type", "GateWay")

	err := RegisterSever(conf, nil, "gate1")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		node, errs := SelectService("GateWay")
		if errs != nil {
			t.Fatal(errs)
		}
		t.Log("found:", node)
		fmt.Println("node:", node)
		time.Sleep(time.Second)
	}

}

func TestListServer(t *testing.T) {
	conf := viper.New()
	conf.Set("game11.listenAddr", "0.0.0.0:3011")
	conf.Set("game11.type", "Game")
	conf.Set("game12.listenAddr", "0.0.0.0:3012")
	conf.Set("game12.type", "Game")
	err := RegisterSever(conf, nil, "game11", "game12")
	if err != nil {
		t.Fatal(err)
	}

	nodes := ListService("GateWay")
	t.Logf("found:%+v", len(nodes))
	for _, v := range nodes {
		t.Log("found node:", v)
	}
}
