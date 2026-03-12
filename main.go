package main

import (
	_ "x-admin/internal/packed"

	_ "x-admin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"x-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
