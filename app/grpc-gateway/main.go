package main

import (
	"proxima/app/grpc-gateway/internal/cmd"
	_ "proxima/app/grpc-gateway/internal/packed"
	"proxima/pkg/boot"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var ctx = gctx.GetInitCtx()
	boot.Init()
	cmd.Main.Run(ctx)
}
