package cmd

import (
	"context"
	"fmt"
	"net/http"

	"proxima/app/grpc-gateway/internal/gateway"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start grpc-gateway server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initCtx := gctx.GetInitCtx()
			address := g.Cfg().MustGet(initCtx, "server.address", ":8100").String()

			mux, err := gateway.NewServeMux(ctx)
			if err != nil {
				return err
			}

			handler := gateway.WithMiddleware(mux)

			g.Log().Infof(ctx, "grpc-gateway server starting on %s", address)
			if err = http.ListenAndServe(address, handler); err != nil {
				return fmt.Errorf("failed to serve: %w", err)
			}
			return nil
		},
	}
)
