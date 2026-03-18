package main

import (
	_ "proxima/app/user/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/user/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
