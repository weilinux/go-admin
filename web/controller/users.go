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
	"strconv"
	"time"
)

func UserLogin(c *gin.Context) {
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

func UserLogout(c *gin.Context) {
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

func UserSignup(c *gin.Context) {
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

func DeleteUser(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	// TODO: 这个删除不一定成功执行的, 需要后面的接口返回错误进行检查
	_ = model.DeleteUser(ID)
	response.ToResponse(SuccessResponse{
		Msg: "删除用户成功",
	})
}

func EditUser(c *gin.Context) {
	response := NewResponse(c)
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := &model.User{Model: &model.Model{ID: uint(ID)}}

	user.UserName = c.PostForm("UserName")
	user.Password = c.PostForm("Password")

	userDetails, db := model.GetUserById(ID)
	if user.UserName != "" {
		userDetails.UserName = user.UserName
	}
	if user.Password != "" {
		userDetails.Password = user.Password
	}

	// TODO: 修改用户成功后, update time没有更新!
	db.Save(&userDetails)
	response.ToResponse(SuccessResponse{
		Msg:  "修改用户配置成功",
		Data: userDetails,
	})
}

func AddUser(c *gin.Context) {
	response := NewResponse(c)
	var user = &model.User{}
	user.UserName = c.PostForm("UserName")
	user.Password = c.PostForm("Password")
	ID := c.PostForm("ID")
	id, _ := strconv.Atoi(ID)

	if model.UserExists(*user) {
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	_, err := model.CreateUser(model.User{
		Model: &model.Model{
			ID:          uint(id),
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
		UserName: user.UserName,
		Password: string(generatedHash([]byte(user.Password))),
	})
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(SuccessResponse{
		Msg:  "添加用户成功",
		Data: user,
	})
}

func UserInfo(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
		return
	}
	user, _ := model.GetUserById(ID)
	response.ToResponse(SuccessResponse{Data: user})
}

func GetUsers(c *gin.Context) {
	response := NewResponse(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	rows, count := model.Paginate(&model.User{}, page, limit)

	data := map[string]interface{}{
		"message": "用户管理界面",
		"rows":    rows,
		"total":   count,
	}

	response.ToResponse(SuccessResponse{
		Data: data,
		Code: 200,
	})
}

// TODO: 获取个人信息完善
func GetUserProfile(c *gin.Context) {
	response := NewResponse(c)
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUserProfile(id)
	response.ToResponse(SuccessResponse{Data: data, Code: code})
}

// Optimi 优化点,更新记录的方式
// var data map[string]string
//
// if err := c.BodyParser(&data); err != nil {
// return err
// }
//
// cookie := c.Cookies("jwt")
//
// id, _ := util.ParseJwt(cookie)
//
// userId, _ := strconv.Atoi(id)
//
// user := models.User{
// Id:        uint(userId),
// FirstName: data["first_name"],
// LastName:  data["last_name"],
// Email:     data["email"],
// }
//
// database.DB.Model(&user).Updates(user)

// TODO: 更新个人信息完善
func UpdateUserProfile(c *gin.Context) {
	response := NewResponse(c)
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateUserProfile(id, &data)

	response.ToResponse(SuccessResponse{Code: code})
}

// TODO: 修改用户密码
func ChangeUserPassword(c *gin.Context) {
	response := NewResponse(c)
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassword(id, &data)

	response.ToResponse(SuccessResponse{Code: code})

	// TODO: 修改用户密码:github td27-admin
	// var mp systemReq.ModifyPass
	// _ = c.ShouldBindJSON(&mp)
	//
	// // 参数校验
	// validate := validator.New()
	// if err := validate.Struct(&mp); err != nil {
	// 	response.FailWithMessage("请求参数错误", c)
	// 	global.TD27_LOG.Error("请求参数错误", zap.Error(err))
	// 	return
	// }
	//
	// if err := userService.ModifyPass(mp); err != nil {
	// 	response.FailWithMessage("修改失败", c)
	// 	global.TD27_LOG.Error("修改失败", zap.Error(err))
	// } else {
	// 	response.OkWithMessage("修改成功", c)
	// }
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
