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
	if err != nil {
		cookie = uuid.NewV4().String()
	}
	c.SetCookie("session2", cookie, 3600, "/", "www.wllinux.com", false, true)

	u := DbSessions[cookie]

	if _, err := model.FindUserByName(u); err != nil {
		return ""
	}
	return u
}

func SetCookieLogin(c *gin.Context) {
	cookie := uuid.NewV4().String()
	// c.SetCookie("session2", cookie, 3600, "/", "www.wllinux.com", false, true)
	DbSessions[cookie] = c.PostForm("UserName")
	c.SetCookie("session", cookie, 3600, "/", "www.wllinux.com", false, true)
	return
}
