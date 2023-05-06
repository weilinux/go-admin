package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
	"github.com/weilinux/go-gin-skeleton-auth/app"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

func RequestLog() gin.HandlerFunc {
	skip := map[string]int{
		"/":       1,
		"/status": 1,
		"/health": 1,
	}

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		path := c.Request.URL.Path
		reqId := strutil.Md5(fmt.Sprintf("%d", start.Nanosecond()))

		// add reqId to context
		c.Set("reqId", reqId)

		// c.MustBindWith()
		// Process request
		c.Next()

		if _, ok := skip[path]; ok {
			return
		}

		// log post/put data
		postData := ""
		if c.Request.Method != "GET" {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			postData = string(buf)
		}

		app.Logger.Info(
			"complete request",
			zap.String("req_id", reqId),
			zap.Namespace("context"),
			zap.String("req_date", start.Format(app.DateFormat)),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.URL.String()),
			zap.String("client_ip", c.ClientIP()),
			zap.Int("http_status", c.Writer.Status()),
			zap.String("elapsed_time", mathutil.ElapsedTime(start)),
			zap.String("post_data", postData),
		)
	}
}
