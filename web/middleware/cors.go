package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Access-Control-Allow-Credentials=true和Access-Control-Allow-Origin="*"有冲突
		// 故Access-Control-Allow-Origin需要指定具体得跨域origin
		c.Header("Access-Control-Allow-Origin", "http://localhost:8084")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Header("Access-Control-Allow-Headers", "content-type")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		// c.Header("Access-Control-Expose-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "")
			c.Abort()
			return
		}
		c.Next()
	}
}

// 在 middleware 中使用
// router.Use(cors.New(auth.CorsConfig(config)))
//
// func CorsConfig(conf *config.Configuration) cors.Config {
//
// 	corsConf := cors.Config{
// 		MaxAge:                 12 * time.Hour,
// 		AllowBrowserExtensions: true,
// 	}
//
// 	if mode.IsDev() {
// 		// 在開發環境時，允許所有 origins、所有 methods 和多數的 headers
// 		corsConf.AllowAllOrigins = true
// 		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
// 		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
// 			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
// 	} else {
// 		// 在正式環境時則根據設定檔調整
// 		compiledOrigins := compileAllowedCORSOrigins(conf.Server.Cors.AllowOrigins)
// 		corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
// 		corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Origin",
// 			"Connection", "Accept-Encoding", "Accept-Language", "Host"}
// 		corsConf.AllowOrigins = []string{"https://www.example.com"}
// 	}
//
// 	return corsConf
// }
