package main

import (
	"flag"

	_ "net/http/pprof"

	"github.com/gin-gonic/gin"

	_ "github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/router"
)

func main() {
	var port, mode string
	flag.StringVar(&port, "port", "3000", "service listening at, default 3000")
	flag.StringVar(&mode, "mode", "debug", "service running mode, default debug mode")

	flag.Parse()

	gin.SetMode(mode)
	router.Init(port)
}
