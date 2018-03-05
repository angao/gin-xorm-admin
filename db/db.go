package db

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

// X 全局DB
var x *xorm.Engine

func init() {
	var err error
	cfg, err := ini.Load("config/db.ini")
	if err != nil {
		log.Fatal(err)
	}

	username := cfg.Section("mysql").Key("username").Value()
	password := cfg.Section("mysql").Key("password").Value()
	url := cfg.Section("mysql").Key("url").Value()

	source := fmt.Sprintf("%s:%s@%s", username, password, url)
	x, err = xorm.NewEngine("mysql", source)

	if err != nil {
		log.Fatal(err)
	}

	err = x.RegisterSqlMap(xorm.Xml("./db/sql/xml", ".xml"))
	if err != nil {
		log.Fatal(err)
	}

	err = x.RegisterSqlTemplate(xorm.Default("./db/sql/tpl", ".sql"))
	if err != nil {
		log.Fatal(err)
	}

	err = x.StartFSWatcher()
	if err != nil {
		log.Printf("sql parse error: %#v\n", err)
	}

	// 30minute ping db to keep connection
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			x.Ping()
		}
	}(x)
	//  x.ShowSQL(true)
}
