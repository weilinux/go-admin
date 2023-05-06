package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/web/middleware"
	"github.com/weilinux/go-gin-skeleton-auth/web/session"
	"net/http"
	"strconv"
	"time"
)

// UserLogin
// # Step 1: Add a ```login.gohtml``` page.
// # Step 2: Handle the submission of a login form
// # Step 3: A fake user was added to dbUsers in func init
// # Step 4: md5 Password
// # STEP 5: cookie mechanism Session ID
func UserLogin(c *gin.Context) {
	var userInput model.User
	userInput.UserName = c.PostForm("UserName")
	userInput.Password = c.PostForm("Password")

	if UserAuth(userInput) {
		sID := uuid.NewV4()
		c.SetCookie("session", sID.String(), 3600, "/", "localhost", false, true)
		session.DbSessions[sID.String()] = c.PostForm("UserName")

		token, err := generateJWT(userInput.UserName)
		if err != nil {
			fmt.Println("error generating token")
		}
		c.JSON(http.StatusOK, gin.H{
			"UserName": userInput.UserName,
			"Password": userInput.Password,
			"Token":    token,
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"UserName": userInput.UserName,
		})
		return
	}
}

func UserLogout(c *gin.Context) {
	// authenticate.ExpireUserSession(w, r)
	// authenticate.ExpireSecureCookie(w, r)

	cookie, err := c.Cookie("session")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	delete(session.DbSessions, cookie)
	c.SetCookie("session", "", -1, "/", "localhost", false, true)

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

// TODO: add jwt authentication

func UserSignup(c *gin.Context) {
	var userInput model.User
	userInput.UserName = c.PostForm("UserName")
	userInput.Password = c.PostForm("Password")

	if model.UserExists(userInput) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "user already exists",
		})
		return
	}

	if _, err := model.CreateUser(model.User{
		UserName: userInput.UserName,
		Password: string(generatedHash([]byte(userInput.Password))),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Insert user failed",
		})
	}

	c.Redirect(http.StatusSeeOther, "/")
	return

}

// different between session and cookie
// 1 session server side; cookie client side
// 2 session dependent on Cookie, but Cookie is not dependent on session
// 3 session end where browser closed; cookie expires by lifetime you set
// 4 during session you can store as much as you want, but max size for cookie is 4KB
// 5 you can use session functions to deal with session, but cookie is not

// session
// session是服务端存储的一个全局变量，每个session有一个独特的ID，用来取值。
// 当一个session创建的时候，a cookie(包含了session id)会保存在客户端，并且跟请求一起发送到服务端。
// 如果客户端不支持Cookie, 那这个unique session id 会显示在URL中。
// The session values are automatically deleted when the browser is closed. If you
// want to store the values permanently, then you should store them in the database.

// cookie
// Once a cookie has been set, all page requests that follow return the cookie name and value
// cookie 不能跨域  A cookie can only be read from the domain that it has been issued from

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if ID, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		_ = model.DeleteUser(ID)
	}
}

func EditUser(c *gin.Context) {
	var user = &model.User{}

	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user.ID = uint(ID)

	user.UserName = c.PostForm("UserName")
	user.Password = c.PostForm("Password")

	userDetails, db := model.GetUserById(ID)
	if user.UserName != "" {
		userDetails.UserName = user.UserName
	}
	if user.Password != "" {
		userDetails.Password = user.Password
	}
	db.Save(&userDetails)
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated",
		"user":    userDetails,
	})
	return
}

func AddUser(c *gin.Context) {
	var user = &model.User{}
	user.UserName = c.PostForm("UserName")
	user.Password = c.PostForm("Password")

	if model.UserExists(*user) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "user already exists",
		})
		return
	}

	if _, err := model.CreateUser(model.User{
		UserName: user.UserName,
		Password: string(generatedHash([]byte(user.Password))),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Insert user failed",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// UserInfo TODO: implement 这里需要再做分页
func UserInfo(c *gin.Context) {
	id := c.Param("id")
	if ID, err := strconv.ParseInt(id, 10, 64); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		user, _ := model.GetUserById(ID)
		data := map[string]interface{}{
			"user": user,
		}
		c.JSON(http.StatusOK, data)
	}
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
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(middleware.SampleSecretKey)

	if err != nil {
		fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
