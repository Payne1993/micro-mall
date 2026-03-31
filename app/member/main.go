package main

import (
	_ "proxima/app/member/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/member/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
