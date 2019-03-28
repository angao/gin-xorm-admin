package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/router"
)

var (
	port, mode string
)

func init() {
	flag.StringVar(&port, "port", "3000", "server listening on, default 3000")
	flag.StringVar(&mode, "mode", "debug", "server running mode, default debug mode")
}

func main() {
	flag.Parse()

	gin.SetMode(mode)
	router := router.Init()
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Start server: %+v", err)
	}
}
