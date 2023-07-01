package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"github.com/weilinux/go-gin-skeleton-auth/web/session"
	"net/http"
	"strconv"
)

type hostAssignment struct {
	UserId  int64   `json:"user_id" binding:"required"`
	HostIds []int64 `json:"host_id" binding:"required"`
}

func GetBindHosts(c *gin.Context) {
	response := NewResponse(c)
	// old sytle get username
	// user, _ := model.FindUserByName(session.GetUser(c))

	// 20230505
	var username string
	if user, ok := c.Get("username"); ok == true {
		username = user.(string)
	}
	user, err := model.FindUserByName(username)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	rows, count := model.GetUserHosts(user.ID, page, limit)
	data := map[string]interface{}{
		"status":   "UP",
		"rows":     rows,
		"total":    count,
		"UserName": session.GetUser(c),
	}
	response.ToResponse(SuccessResponse{
		Data: data,
	})
}

func GetUnBindHosts(c *gin.Context) {
	response := NewResponse(c)
	// user, _ := model.FindUserByName(session.GetUser(c))
	// username, _ := c.Get("username")
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
	}
	data := map[string]interface{}{
		"status":   "UP",
		"rows":     model.GetUserUnbindHosts(ID),
		"UserName": session.GetUser(c),
	}
	response.ToResponse(SuccessResponse{Data: data})
}

func DeleteHost(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
	}
	_ = model.DeleteHost(ID)
	response.ToResponse(SuccessResponse{Msg: "删除主机成功"})
}

func EditHost(c *gin.Context) {
	response := NewResponse(c)
	var host = &model.Host{}

	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	host.ID = uint(ID)

	host.HostName = c.PostForm("HostName")
	host.HostIP = c.PostForm("HostIP")
	host.HostPort, _ = strconv.Atoi(c.PostForm("HostPort"))

	hostDetails, db := model.GetHostById(ID)
	if host.HostName != "" {
		hostDetails.HostName = host.HostName
	}
	if host.HostIP != "" {
		hostDetails.HostIP = host.HostIP
	}
	if host.HostPort != 0 {
		hostDetails.HostPort = host.HostPort
	}

	// model.DBM.Model(&model.Host{}).Where("id= ?", id).Update(&hostDetails)
	db.Save(&hostDetails)

	response.ToResponse(SuccessResponse{
		Msg:  "修改主机成功",
		Data: hostDetails,
	})
}

func HostInfo(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
	}
	host, _ := model.GetHostById(ID)
	response.ToResponse(SuccessResponse{
		Code: 0,
		Data: host,
	})
}

func AddHost(c *gin.Context) {
	response := NewResponse(c)
	var host model.Host
	host.HostName = c.PostForm("HostName")
	host.HostIP = c.PostForm("HostIP")
	host.HostPort, _ = strconv.Atoi(c.PostForm("HostPort"))
	host.CreateHost()
	response.ToResponse(SuccessResponse{
		Msg:  "添加主机成功",
		Data: host,
	})
}

func AssignHost(c *gin.Context) {
	response := NewResponse(c)

	var ha hostAssignment
	if err := c.ShouldBindJSON(&ha); err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
	}
	userId := ha.UserId
	hostIds := ha.HostIds

	err := model.HostAssignment(userId, hostIds)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
	}
	response.ToResponse(SuccessResponse{Msg: "分配主机成功"})
}

func SshHost(c *gin.Context) {
	// TODO: url主机地址加密 https://cloud.tencent.com/developer/article/1469183
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	host, _ := model.GetHostById(ID)
	if !host.IsAlive() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Can't not connect to host",
		})

	}

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	sshAccessPath := "http://" + host.HostIP + ":3003/wetty"
	c.Redirect(http.StatusFound, sshAccessPath)
}

// https://github.com/yohanapriyandi/Golang-Rest-API-With-GIN-GORM.git
