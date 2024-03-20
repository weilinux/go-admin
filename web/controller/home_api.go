package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/app"
	"github.com/weilinux/go-gin-skeleton-auth/web/session"
	"net/http"
)

func Home(c *gin.Context) {
	u := session.GetUser(c)
	c.JSON(200, gin.H{
		"UserName": u,
	})
}

// func SwagDoc(c *gin.Context) {
// 	fInfo, _ := os.Stat("static/swagger.json")
//
// 	data := map[string]string{
// 		"EnvName":    app.EnvName,
// 		"AppName":    app.Name,
// 		"JsonFile":   "/static/swagger.json",
// 		"SwgUIPath":  "/static/swagger-ui",
// 		"AssetPath":  "/static",
// 		"UpdateTime": fInfo.ModTime().Format(app.DateFormat),
// 	}
//
// 	c.HTML(200, "swagger.tpl", data)
//
// }

// func SwagDoc(_ *gin.Context) gin.HandlerFunc {
// 	return ginSwagger.WrapHandler(swaggerfiles.Handler)
// }

// @Tags InternalApi
// @Summary 检测API
// @Description get app health
// @Success 201 {string} json data
// @Failure 403 body is empty
// @Router /health [get]
func AppHealth(c *gin.Context) {
	data := map[string]interface{}{
		"status": "HEALTH",
		"info":   app.GitInfo,
	}
	c.JSON(http.StatusOK, data)
}

// @Tags InternalApi
// @Summary 状态API
// @Description get app status
// @Success 201 {string} json data
// @Failure 403 body is empty
// @Router /status [get]
func AppStatus(c *gin.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}
	c.JSON(http.StatusOK, data)
}
