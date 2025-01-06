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

	authApi := new(controller.AuthApi)

	index := noAuth.Group("/")
	{
		index.GET("/", controller.Home)
		index.GET("/health", controller.AppHealth)
		index.GET("/status", controller.AppStatus)
		index.POST("/login", authApi.UserLogin)
		index.POST("/signup", authApi.UserSignup)
		// 终端管理
		index.GET("/ws", controller.ShellWs)
	}

	// admin for user login
	admin := r.Group("/api/v1")
	admin.Use(middleware.Auth())
	admin.Use(middleware.Sensitive())
	admin.Use(middleware.Cors())

	admin.GET("/logout", authApi.UserLogout)
	// 用户管理
	user := admin.Group("/")
	{
		userApi := new(controller.UserApi)
		user.GET("/users", userApi.GetUsers)
		user.DELETE("/users/:id", userApi.DeleteUser)
		user.PUT("/users/:id", userApi.EditUser)
		user.GET("/users/:id", userApi.UserInfo)
		user.POST("/users", userApi.AddUser)
		user.GET("user/profile/:id", userApi.GetUserProfile)
		user.PUT("user/profile/:id", userApi.UpdateUserProfile)
		user.PUT("user/changepw/:id", userApi.ChangeUserPassword)
	}
	// 角色管理
	roleApi := new(controller.RoleApi)
	role := admin.Group("/")
	{
		role.GET("/roles", roleApi.GetRoles)
		role.DELETE("/roles/:id", roleApi.DeleteRole)
		role.PUT("/roles/:id", roleApi.EditRole)
		role.GET("/roles/:id", roleApi.RoleInfo)
		role.POST("/roles", roleApi.AddRole)
	}

	// 权限管理
	permApi := new(controller.PermissionApi)
	permission := admin.Group("/")
	{
		permission.GET("/permissions", permApi.GetPermissions)
	}

	// 主机管理
	hostApi := new(controller.HostApi)
	host := admin.Group("/")
	{
		host.GET("/users/hosts", hostApi.GetBindHosts)
		host.GET("/users/searchhost", hostApi.SearchHosts)
		host.GET("/users/:id/unbindhosts", hostApi.GetUnBindHosts)
		host.POST("/hosts/assign", hostApi.AssignHost)
		host.POST("/users/:id/hosts", hostApi.AssignHost)

		host.POST("/hosts", hostApi.AddHost)
		host.PUT("/hosts/:id", hostApi.EditHost)
		host.GET("/hosts/:id", hostApi.HostInfo)
		host.DELETE("/hosts/:id", hostApi.DeleteHost)
	}

	// TODO: add 新开tab的标题应该是服务器的主机名称
	xterm := admin.Group("/")
	{
		xterm.GET("/hosts/:id/ssh", hostApi.SshHost)
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
