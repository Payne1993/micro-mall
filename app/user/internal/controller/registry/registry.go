package registry

import (
	"sync"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type RegisterFunc func(*grpcx.GrpcServer)

var (
	mu            sync.RWMutex
	registerFuncs []RegisterFunc
)

func Add(fn RegisterFunc) {
	mu.Lock()
	defer mu.Unlock()
	registerFuncs = append(registerFuncs, fn)
}

func RegisterAll(server *grpcx.GrpcServer) {
	mu.RLock()
	funcs := append([]RegisterFunc(nil), registerFuncs...)
	mu.RUnlock()

	for _, fn := range funcs {
		fn(server)
	}
}
