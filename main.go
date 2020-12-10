package main

import (
	_ "go-gf-demo/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
	// <-time.After(time.Second * 100)
}
