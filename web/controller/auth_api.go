package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"github.com/weilinux/go-gin-skeleton-auth/web/middleware"
	"github.com/weilinux/go-gin-skeleton-auth/web/session"
	"net/http"
	"time"
)

type AuthApi struct {
	BaseApi
}

// 用户登录/注销/找回密码

func (auth *AuthApi) UserLogin(c *gin.Context) {
	response := NewResponse(c)
	var userInput model.User
	userInput.UserName = c.PostForm("UserName")
	userInput.Password = c.PostForm("Password")

	if UserAuth(userInput) {
		session.SetCookieLogin(c)
		token, err := generateJWT(userInput.UserName)
		if err != nil {
			response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		}
		data := map[string]interface{}{
			"UserName": userInput.UserName,
			"Token":    token,
		}
		response.ToResponse(SuccessResponse{Data: data})
	} else {
		response.ToErrorResponse(errcode.Fail.WithDetails("用户认证失败!!!"))
	}
}

func (auth *AuthApi) UserLogout(c *gin.Context) {
	// authenticate.ExpireUserSession(w, r)
	// authenticate.ExpireSecureCookie(w, r)

	cookie, err := c.Cookie("session")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	delete(session.DbSessions, cookie)
	c.SetCookie("session", "", -1, "/", "www.wllinux.com", false, true)

	c.Redirect(http.StatusSeeOther, "/login")
}

// func UserLogout(c *gin.Context) {
// 	// authenticate.ExpireUserSession(w, r)
// 	// authenticate.ExpireSecureCookie(w, r)
//
// 	cookie, err := c.Cookie("session")
// 	if err != nil {
// 		c.Redirect(http.StatusSeeOther, "/login")
// 		return
// 	}
// 	delete(session.DbSessions, cookie)
// 	c.SetCookie("session", "", -1, "/", "192.168.2.230", false, true)
//
// 	c.Redirect(http.StatusSeeOther, "/login")
// }

func (auth *AuthApi) UserSignup(c *gin.Context) {
	response := NewResponse(c)
	var userInput model.User
	userInput.UserName = c.PostForm("UserName")
	userInput.Password = c.PostForm("Password")

	if model.UserExists(userInput) {
		response.ToErrorResponse(errcode.ErrorExistUserFail)
		return
	}

	_, err := model.CreateUser(model.User{
		Model: &model.Model{
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
		UserName: userInput.UserName,
		Password: string(generatedHash([]byte(userInput.Password))),
	})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
	}
	response.ToResponse(SuccessResponse{Code: 200})
}

func UserAuth(userInput model.User) bool {
	var (
		user model.User
		err  error
	)
	if user, err = model.FindUserByName(userInput.UserName); err != nil {
		return false
	}
	return compareHash([]byte(user.Password), []byte(userInput.Password))
}

func generateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(middleware.SampleSecretKey)

	if err != nil {
		_ = fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
