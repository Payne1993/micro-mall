package main

import (
	_ "proxima/app/gateway/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/gateway/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
