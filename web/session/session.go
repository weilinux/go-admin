package session

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/weilinux/go-gin-skeleton-auth/model"
)

var DbSessions = map[string]string{}

func GetUser(c *gin.Context) string {
	// get cookie
	cookie, err := c.Cookie("session")
	// 这时为什么要生产一个session 保存至cookie中呢 ?
	// 这时为什么要生产一个session 保存至cookie中呢 ?
	// 这时为什么要生产一个session 保存至cookie中呢 ?
	if err != nil {
		cookie = uuid.NewV4().String()
	}
	c.SetCookie("session", cookie, -1, "/", "192.168.2.106", false, true)

	u := DbSessions[cookie]

	if _, err := model.FindUserByName(u); err != nil {
		return ""
	}
	return u
}
