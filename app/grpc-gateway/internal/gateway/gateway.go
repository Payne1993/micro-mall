package gateway

import (
	"context"
	"fmt"
	"net/http"

	_ "proxima/app/grpc-gateway/internal/handlers" // auto-register all handlers

	"proxima/app/grpc-gateway/internal/middleware"
	"proxima/app/grpc-gateway/internal/registry"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	gwRuntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

type routeCfg struct {
	Name       string   `json:"name"`
	Predicates []string `json:"predicates"`
}

// NewServeMux creates and configures a grpc-gateway ServeMux with all upstream handlers registered.
func NewServeMux(ctx context.Context) (*gwRuntime.ServeMux, error) {
	mux := gwRuntime.NewServeMux(
		gwRuntime.WithMetadata(forwardMetadata),
		gwRuntime.WithMarshalerOption(gwRuntime.MIMEWildcard, &gwRuntime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		gwRuntime.WithIncomingHeaderMatcher(incomingHeaderMatcher),
	)

	initCtx := gctx.GetInitCtx()
	var routes []routeCfg
	if err := g.Cfg().MustGet(initCtx, "gateway.routes").Scan(&routes); err != nil {
		return nil, fmt.Errorf("failed to load gateway.routes config: %w", err)
	}

	for _, route := range routes {
		handlers := registry.Get(route.Name)
		if len(handlers) == 0 {
			g.Log().Warningf(ctx, "no handlers registered for service %q, skipping", route.Name)
			continue
		}

		conn, err := grpcx.Client.NewGrpcClientConn(route.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to service %q: %w", route.Name, err)
		}

		for _, fn := range handlers {
			if err = fn(context.Background(), mux, conn); err != nil {
				return nil, fmt.Errorf("failed to register handler for service %q: %w", route.Name, err)
			}
		}

		g.Log().Infof(ctx, "registered %d handler(s) for service %q, predicates: %v",
			len(handlers), route.Name, route.Predicates)
	}

	return mux, nil
}

// WithMiddleware wraps the grpc-gateway ServeMux with HTTP middleware and a health endpoint.
func WithMiddleware(gwMux *gwRuntime.ServeMux) http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint (bypasses grpc-gateway).
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// All other requests are handled by grpc-gateway.
	mux.Handle("/", gwMux)

	// Apply middleware chain (outermost runs first).
	var handler http.Handler = mux
	handler = middleware.RateLimiter(handler)
	handler = middleware.CORS(handler)
	return handler
}

// forwardMetadata converts incoming HTTP headers to gRPC metadata.
func forwardMetadata(ctx context.Context, req *http.Request) metadata.MD {
	md := metadata.MD{}
	for key, values := range req.Header {
		for _, value := range values {
			md.Append(key, value)
		}
	}
	if host := req.Host; host != "" {
		md.Set("x-forwarded-host", host)
	}
	if requestID := req.Header.Get("X-Request-Id"); requestID != "" {
		md.Set("x-request-id", requestID)
	}
	return md
}

// incomingHeaderMatcher allows specific HTTP headers to be forwarded as gRPC metadata.
func incomingHeaderMatcher(key string) (string, bool) {
	switch key {
	case "Authorization", "X-Request-Id", "X-Forwarded-For", "X-Real-Ip":
		return key, true
	default:
		return gwRuntime.DefaultHeaderMatcher(key)
	}
}
