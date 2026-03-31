package main

import (
	"micro-mall.dev/app/grpc-gateway/internal/cmd"
	_ "micro-mall.dev/app/grpc-gateway/internal/packed"
	"micro-mall.dev/pkg/boot"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var ctx = gctx.GetInitCtx()
	boot.Init()
	cmd.Main.Run(ctx)
}
