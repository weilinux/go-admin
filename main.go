package main

import (
	"embed"
	"github.com/gookit/color"
	"github.com/weilinux/go-gin-skeleton-auth/app"
	"github.com/weilinux/go-gin-skeleton-auth/model/myrds"
	"github.com/weilinux/go-gin-skeleton-auth/model/mysql"
	"github.com/weilinux/go-gin-skeleton-auth/web"
	"log"
	"os"
)

// staticFS is an embedded file system
// //go:embed public/*
var staticFS embed.FS

func init() {
	var err error
	app.BootStrap("./config")

	// - redis, mongo, mysql connection
	err = myrds.InitRedis()
	checkError("Rds init error:", err)

	// err = mongo.InitMongo()
	// checkError("Mgo init error:", err)

	web.InitServer(staticFS)
}

func main() {
	sqlDB := mysql.ConnectDatabase()

	// init services
	log.Printf("================ Begin Running(PID: %d) ================", os.Getpid())

	web.Run()
	defer myrds.ClosePool()
	defer mysql.CloseDatabase(sqlDB)
}

func checkError(prefix string, err error) {
	if err != nil {
		color.Error.Println(prefix, err.Error())
		os.Exit(2)
	}
}
