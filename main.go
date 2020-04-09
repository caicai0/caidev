package main

import (
	_ "caidev/boot"
	_ "caidev/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
