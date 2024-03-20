package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/web/middleware"
	"github.com/weilinux/go-gin-skeleton-auth/web/session"
)

// JsonMapData only use for swagger docs
type JsonMapData struct {
	Code    int               `json:"code" description:"return code, 0 is successful."`
	Message string            `json:"message" description:"return message"`
	Data    map[string]string `json:"data" description:"return data"`
}

// BaseApi controller
type BaseApi struct {
	lang string
}

type Controller struct {
}

type Message struct {
	Type    string // success, warning, error, etc...
	Content string
}

type PageData struct {
	Title           string
	Messages        []Message
	IsAuthenticated bool
}

func isAuthenticated(c *gin.Context) bool {
	cookie, err := c.Cookie(middleware.UserIDKey)
	if err != nil {
		return false
	}
	username := session.DbSessions[cookie]
	if _, err = model.FindUserByName(username); err != nil {
		return false
	}
	return true
}

func (controller Controller) DefaultPageData(c *gin.Context) PageData {
	return PageData{
		Title:           "Home",
		Messages:        nil,
		IsAuthenticated: isAuthenticated(c),
	}
}
