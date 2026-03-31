package main

import (
	_ "micro-mall.dev/app/word/internal/packed"
	"micro-mall.dev/pkg/boot"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"micro-mall.dev/app/word/internal/cmd"
)

func main() {
	var ctx = gctx.New()
	boot.Init()
	cmd.Main.Run(ctx)
}
