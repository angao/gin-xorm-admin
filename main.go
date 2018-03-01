package main

import (
	_ "github.com/angao/gin-xorm-admin/db"
	"github.com/angao/gin-xorm-admin/router"
)

func main() {
	router.Init()
}
