package main

import (
	_ "micro-mall.dev/app/member/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/joho/godotenv"

	"micro-mall.dev/app/member/internal/cmd"
)

func main() {
	_ = godotenv.Load()
	cmd.Main.Run(gctx.GetInitCtx())
}
