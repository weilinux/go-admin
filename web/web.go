package web

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"github.com/weilinux/go-gin-skeleton-auth/app"
	"github.com/weilinux/go-gin-skeleton-auth/web/middleware"
	"github.com/weilinux/go-gin-skeleton-auth/web/router"
	"io/fs"
	"log"
	"net/http"
)

var server *gin.Engine

// func Server() *gin.Engine {
// 	return server
// }

func InitServer(staticFS embed.FS) {
	server = gin.New()

	if app.IsEnv(app.EnvDev) {
		server.Use(gin.Logger(), gin.Recovery())
	}

	// 生产环境的日志配置
	server.Use(middleware.RequestLog())

	// v1 := server.Group("/v1")
	// We parse and load the html files into our t variable
	t, err := loadTemplates(staticFS)
	if err != nil {
		log.Fatalln(err)
	}

	server.SetHTMLTemplate(t)

	// our assets are only located in a section of our file system. so we create a sub file system.
	subFS, err := fs.Sub(staticFS, "public/assets")
	if err != nil {
		log.Fatalln(err)
	}

	// All static assets are under the /assets path, so we make this its own group called assets
	assets := server.Group("/assets")
	// assets := server.Group("/")

	// This middleware sets the Cache-Control header and is applied to the assets group only
	// assets.Use(middleware.Cache(conf.CacheMaxAge))

	// All requests to /assets will use the sub fil system which contains all our static assets
	assets.StaticFS("/", http.FS(subFS))

	// Session middleware is applied to all groups after this point.
	// r.Use(middleware.Session(db))

	router.AddRoutes(server)
}

func Run() {
	err := server.Run(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
	if err != nil {
		color.Error.Println(err)
	}
}
