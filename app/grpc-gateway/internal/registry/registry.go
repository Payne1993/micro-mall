package registry

import (
	"context"
	"sync"

	gwRuntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// HandlerRegisterFunc is a grpc-gateway handler registration function.
type HandlerRegisterFunc func(ctx context.Context, mux *gwRuntime.ServeMux, conn *grpc.ClientConn) error

var (
	mu       sync.RWMutex
	handlers = make(map[string][]HandlerRegisterFunc)
)

// Add registers a grpc-gateway handler function for the given service name.
func Add(serviceName string, fn HandlerRegisterFunc) {
	mu.Lock()
	defer mu.Unlock()
	handlers[serviceName] = append(handlers[serviceName], fn)
}

// Get returns all registered handler functions for the given service name.
func Get(serviceName string) []HandlerRegisterFunc {
	mu.RLock()
	defer mu.RUnlock()
	return handlers[serviceName]
}