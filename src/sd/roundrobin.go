package sd

import (
	"sync"
	"sync/atomic"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
)

var rrIndex uint64 = 0

// RoundRobin is a roundrobin strategy algorithm for node selection
func RoundRobin(services []*registry.Service) selector.Next {
	nodes := make([]*registry.Node, 0, len(services))

	for _, service := range services {
		nodes = append(nodes, service.Nodes...)
	}

	//var i = rand.Int()
	var mtx sync.Mutex

	return func() (*registry.Node, error) {
		mtx.Lock()

		if len(nodes) == 0 {
			return nil, selector.ErrNoneAvailable
		}

		i := int(atomic.AddUint64(&rrIndex, 1))
		node := nodes[i%len(nodes)]

		mtx.Unlock()

		return node, nil
	}
}
