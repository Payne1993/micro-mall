package main

import (
	"micro-mall.dev/pkg/boot"
	_ "micro-mall.dev/app/user/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"micro-mall.dev/app/user/internal/cmd"
)

func main() {
	var ctx = gctx.GetInitCtx()
	boot.Init()
	cmd.Main.Run(ctx)
}
