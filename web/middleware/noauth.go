package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/web/session"
	"net/http"
)

func NoAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session")
		if err != nil {
			c.Next()
			return
		}
		username := session.DbSessions[cookie]
		if _, err = model.FindUserByName(username); err != nil {
			fmt.Printf("could not find user!, err:%v\n", err)
			c.Next()
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/main")
		c.Abort()
	}
}
