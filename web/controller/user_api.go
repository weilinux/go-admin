package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"strconv"
	"time"
)

type UserApi struct {
	BaseApi
}

func (u *UserApi) DeleteUser(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	//
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

func (u *UserApi) EditUser(c *gin.Context) {
	response := NewResponse(c)
	ID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	user := &model.User{Model: &model.Model{ID: uint(ID)}}
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

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

func (u *UserApi) AddUser(c *gin.Context) {
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

func (u *UserApi) UserInfo(c *gin.Context) {
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

// GetUsers
// @Tags UserApi
// @Summary 获取用户列表
// @Description get paginated user list
// @Security Bearer
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} controller.SuccessResponse{data=map[string]interface{}}
// @Router /api/v1/users [get]
func (u *UserApi) GetUsers(c *gin.Context) {
	// TODO:: 使用 context 控制查询超时
	// https://github.com/weilinux/go-admin/blob/c201cc164aced24718d8dc03071570ae4ce5e0bd/TODO_后端技术体系化#L184
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

// GetUserProfile
// @Tags UserApi
// @Summary 获取用户个人资料
// @Description get user profile by ID
// @Param id path int true "User ID"
// @Success 200 {object} controller.SuccessResponse{data=model.Profile}
// @Router /api/v1/users/{id}/profile [get]
func (u *UserApi) GetUserProfile(c *gin.Context) {
	// TODO: 获取个人信息完善
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

// UpdateUserProfile
// @Tags UserApi
// @Summary 更新用户个人资料
// @Description update user profile
// @Param id path int true "User ID"
// @Param data body model.Profile true "User Profile Data"
// @Success 200 {object} controller.SuccessResponse
// @Router /api/v1/users/{id}/profile [put]
func (u *UserApi) UpdateUserProfile(c *gin.Context) {
	// TODO: 更新个人信息完善
	response := NewResponse(c)
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateUserProfile(id, &data)

	response.ToResponse(SuccessResponse{Code: code})
}

// ChangeUserPassword
// @Tags UserApi
// @Summary 修改用户密码
// @Description change user password
// @Param id path int true "User ID"
// @Param data body model.User true "Password Data"
// @Success 200 {object} controller.SuccessResponse
// @Router /api/v1/users/{id}/password [put]
func (u *UserApi) ChangeUserPassword(c *gin.Context) {
	// TODO: 修改用户密码
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
