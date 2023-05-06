package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/app"
	"net/http"
)

// InternalApi func
type InternalApi struct {
	BaseApi
}

func (a *InternalApi) Config(c *gin.Context) {
	var data interface{}

	key := c.Query("key")
	if key == "" {
		data = app.Config.Data()
	} else {
		data = app.Config.Get(key)
	}

	c.JSON(http.StatusOK, data)
}
