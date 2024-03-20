package router

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
	noAuth.Use(middleware.Sensitive())
	// noAuth.Use(middleware.NoAuth())
	noAuth.Use(middleware.Cors())

	index := noAuth.Group("/")
	{
		index.GET("/", controller.Home)
		index.GET("/health", controller.AppHealth)
		index.GET("/status", controller.AppStatus)
		index.POST("/login", controller.UserLogin)
		index.POST("/signup", controller.UserSignup)
		// 终端管理
		index.GET("/ws", controller.ShellWs)
	}

	// admin for user login
	admin := r.Group("/api/v1")
	admin.Use(middleware.Auth())
	admin.Use(middleware.Sensitive())
	admin.Use(middleware.Cors())

	admin.GET("/logout", controller.UserLogout)
	// 用户管理
	user := admin.Group("/")
	{
		user.GET("/users", controller.GetUsers)
		user.DELETE("/users/:id", controller.DeleteUser)
		user.PUT("/users/:id", controller.EditUser)
		user.GET("/users/:id", controller.UserInfo)
		user.POST("/users", controller.AddUser)
		user.GET("user/profile/:id", controller.GetUserProfile)
		user.PUT("user/profile/:id", controller.UpdateUserProfile)
		user.PUT("user/changepw/:id", controller.ChangeUserPassword)
	}
	// 角色管理
	role := admin.Group("/")
	{
		role.GET("/roles", controller.GetRoles)
		role.DELETE("/roles/:id", controller.DeleteRole)
		role.PUT("/roles/:id", controller.EditRole)
		role.GET("/roles/:id", controller.RoleInfo)
		role.POST("/roles", controller.AddRole)
	}

	// 权限管理
	permission := admin.Group("/")
	{
		permission.GET("/permissions", controller.GetPermissions)
	}

	// 主机管理
	host := admin.Group("/")
	{
		host.GET("/users/hosts", controller.GetBindHosts)
		host.GET("/users/searchhost", controller.SearchHosts)
		host.GET("/users/:id/hosts", controller.GetUnBindHosts)
		host.DELETE("/hosts/:id", controller.DeleteHost)
		host.PUT("/hosts/:id", controller.EditHost)
		host.GET("/hosts/:id", controller.HostInfo)
		host.POST("/hosts", controller.AddHost)
		host.POST("/hosts/assign", controller.AssignHost)

	}

	// TODO: add 新开tab的标题应该是服务器的主机名称
	xterm := admin.Group("/")
	{
		xterm.GET("/hosts/:id/ssh", controller.SshHost)
		// xterm.GET("/host/:id/metrics", controller.MonitorHosts)
		r.GET("/host/metrics", controller.MonitorHosts)
		xterm.GET("/host/metrics", controller.MonitorHosts)
	}

	// 查询配置
	internal := new(controller.InternalApi)
	r.GET("/config", internal.Config)

	// 上传文件
	admin.POST("upload", controller.Upload)

	// 优化找不到页面显示的错误
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, controller.JsonMapData{
			Message: "不好意思，找不到这个页面 !!",
			Data:    map[string]string{},
		})
	})
}
