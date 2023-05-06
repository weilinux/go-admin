package web

import (
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	statics "github.com/weilinux/go-gin-skeleton-auth/static"
	"github.com/weilinux/go-gin-skeleton-auth/web/controller"
	"github.com/weilinux/go-gin-skeleton-auth/web/middleware"
	"net/http"
)

// TODO: add 跨域，jwt认证 https://gitee.com/GoProgect/ginBlog/blob/master/routers/router.go
// TODO: add 跨域，jwt认证 https://gitee.com/GoProgect/ginBlog/blob/master/routers/router.go
// TODO: add 跨域，jwt认证 https://gitee.com/GoProgect/ginBlog/blob/master/routers/router.go

func AddRoutes(r *gin.Engine) {
	statics.SwaggerInfo.BasePath = "/"

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/debug/statsviz/*filepath", func(context *gin.Context) {
		if context.Param("filepath") == "/ws" {
			statsviz.Ws(context.Writer, context.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(context.Writer, context.Request)
	})

	noAuth := r.Group("/")
	// noAuth.Use(middleware.NoAuth())
	noAuth.GET("/", controller.Home)
	// status
	noAuth.GET("/health", controller.AppHealth)
	noAuth.GET("/status", controller.AppStatus)

	noAuth.POST("/login", controller.UserLogin)
	noAuth.POST("/signup", controller.UserSignup)

	// admin for user login
	admin := r.Group("/")
	admin.Use(middleware.Auth())
	admin.Use(middleware.Sensitive())
	// admin.Use(middleware.Cors())
	admin.GET("/logout", controller.UserLogout)

	// 用户管理
	admin.GET("/users", controller.GetUsers)
	admin.DELETE("/users/:id", controller.DeleteUser)
	admin.PUT("/users/:id", controller.EditUser)
	admin.GET("/users/:id", controller.UserInfo)
	admin.POST("/users", controller.AddUser)
	// TODO: implement 权限管理

	admin.GET("/users/hosts", controller.GetBindHosts)
	admin.GET("/users/:id/hosts", controller.GetUnBindHosts)
	admin.DELETE("/hosts/:id", controller.DeleteHost)
	admin.PUT("/hosts/:id", controller.EditHost)
	admin.GET("/hosts/:id", controller.HostInfo)
	admin.POST("/hosts", controller.AddHost)
	admin.POST("/hosts/assign", controller.AssignHost)

	// TODO: add 新开tab的标题应该是服务器的主机名称
	admin.GET("/hosts/:id/ssh", controller.SshHost)
	admin.GET("/host/:id/metrics", controller.MonitorHosts)

	// 查询配置
	internal := new(controller.InternalApi)
	r.GET("/config", internal.Config)

	// 优化找不到页面显示的错误
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, controller.JsonMapData{
			Message: "不好意思，找不到这个页面 !!",
			Data:    map[string]string{},
		})
	})
}
