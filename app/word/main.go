package main

import (
	_ "proxima/app/word/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"proxima/app/word/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
