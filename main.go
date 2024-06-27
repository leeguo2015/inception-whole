package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "inception-whole/internal/packed"

	_ "inception-whole/internal/logic"

	"inception-whole/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
