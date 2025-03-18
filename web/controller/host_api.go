package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	systemReq "github.com/weilinux/go-gin-skeleton-auth/model/system/request"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"github.com/weilinux/go-gin-skeleton-auth/web/session"
	"net/http"
	"strconv"
)

type HostApi struct {
	BaseApi
}

type hostAssignment struct {
	UserId int64 `json:"user_id" binding:"required"`
	// HostIds []int64 `json:"host_id" binding:"required"`
	HostId int64 `json:"host_id" binding:"required"`
}

func getUsernameFromContext(c *gin.Context) (string, bool) {
	if user, exists := c.Get("username"); exists {
		// user, ok := u.(*ms.User)
		v, ok := user.(string)
		return v, ok
	}
	return "", false
}

// GetHosts
// @Tags HostApi
// @Summary 获取主机列表
// @Description get paginated host list
// @Security Bearer
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} controller.SuccessResponse{data=map[string]interface{}}
// @Router /api/v1/hosts [get]
func (h *HostApi) GetHosts(c *gin.Context) {
	response := NewResponse(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	rows, count := model.GetHosts(page, limit)
	data := map[string]interface{}{
		"status": "UP",
		"rows":   rows,
		"total":  count,
	}
	response.ToResponse(SuccessResponse{
		Data: data,
	})
}

// GetBindHosts
// @Tags HostApi
// @Summary 获取用户已绑定主机信息
// @Description get user host info
// @Security Bearer
// @Success 200 {string} json data
// @Failure 403 body is empty
// @Router /api/v1/users/hosts [get]
func (h *HostApi) GetBindHosts(c *gin.Context) {
	response := NewResponse(c)

	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	rows, count := model.GetUserHosts(ID, page, limit)
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

// SearchHosts
// @Tags HostApi
// @Summary 搜索主机
// @Description search hosts by host name
// @Security Bearer
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param host query string false "Host name for filtering"
// @Success 200 {object} controller.SuccessResponse{data=map[string]interface{}}
// @Router /api/v1/hosts/search [get]
func (h *HostApi) SearchHosts(c *gin.Context) {
	response := NewResponse(c)
	// var username string
	// if user, ok := c.Get("username"); ok == true {
	// 	username = user.(string)
	// }
	username, exist := getUsernameFromContext(c)

	if !exist {
		response.ToErrorResponse(errcode.NotFound.WithDetails("username not found in context"))
		return
	}

	user, err := model.FindUserByName(username)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	host := c.DefaultQuery("host", "")

	rows, count := model.GetUserFilterHosts(user.ID, host, page, limit)
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

// GetUnBindHosts
// @Tags HostApi
// @Summary 获取未绑定的主机列表
// @Description get list of hosts not bound to a user
// @Security Bearer
// @Param id path int true "User ID"
// @Success 200 {object} controller.SuccessResponse{data=map[string]interface{}}
// @Router /api/v1/users/{id}/hosts/unbind [get]
func (h *HostApi) GetUnBindHosts(c *gin.Context) {
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

// DeleteHost
// @Tags HostApi
// @Summary 删除主机
// @Description delete a host by ID
// @Security Bearer
// @Param id path int true "Host ID"
// @Success 200 {object} controller.SuccessResponse
// @Router /api/v1/hosts/{id} [delete]
func (h *HostApi) DeleteHost(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
	}
	_ = model.DeleteHost(ID)
	response.ToResponse(SuccessResponse{Msg: "删除主机成功"})
}

// EditHost
// @Tags HostApi
// @Summary 编辑主机信息
// @Description update host information
// @Security Bearer
// @Param id path int true "Host ID"
// @Param hostUpdates body systemReq.EditHost true "Host information to update"
// @Success 200 {object} controller.SuccessResponse{data=model.Host}
// @Router /api/v1/hosts/{id} [put]
func (h *HostApi) EditHost(c *gin.Context) {
	response := NewResponse(c)
	var hostUpdates = &systemReq.EditHost{}

	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := c.ShouldBind(hostUpdates); err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
	}

	hostDetails, db := model.GetHostById(ID)
	if hostUpdates.HostName != "" {
		hostDetails.HostName = hostUpdates.HostName
	}
	if hostUpdates.HostIP != "" {
		hostDetails.HostIP = hostUpdates.HostIP
	}
	if hostUpdates.HostPort != 0 {
		hostDetails.HostPort = hostUpdates.HostPort
	}

	// model.db.Model(&model.Host{}).Where("id= ?", id).Update(&hostDetails)
	db.Save(&hostDetails)

	response.ToResponse(SuccessResponse{
		Msg:  "修改主机成功",
		Data: hostDetails,
	})
}

// HostInfo
// @Tags HostApi
// @Summary 获取主机详情
// @Description get detailed information about a host
// @Security Bearer
// @Param id path int true "Host ID"
// @Success 200 {object} controller.SuccessResponse{data=model.Host}
// @Router /api/v1/hosts/{id} [get]
func (h *HostApi) HostInfo(c *gin.Context) {
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

// AddHost
// @Tags HostApi
// @Summary 添加主机
// @Description add a new host
// @Security Bearer
// @Param HostName formData string true "Host name"
// @Param HostIP formData string true "Host IP address"
// @Param HostPort formData int true "Host port number"
// @Success 200 {object} controller.SuccessResponse{data=model.Host}
// @Router /api/v1/hosts [post]
func (h *HostApi) AddHost(c *gin.Context) {
	response := NewResponse(c)
	var host model.Host
	host.HostName = c.PostForm("HostName")
	host.HostIP = c.PostForm("HostIP")
	host.HostPort, _ = strconv.Atoi(c.PostForm("HostPort"))
	_, err := host.CreateHost()
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
	}
	response.ToResponse(SuccessResponse{
		Msg:  "添加主机成功",
		Data: host,
	})
}

// AssignHost
// @Tags HostApi
// @Summary 分配主机给用户
// @Description assign host to a user
// @Security Bearer
// @Param assignment body controller.hostAssignment true "User and host assignment information"
// @Success 200 {object} controller.SuccessResponse
// @Router /api/v1/hosts/assign [post]
func (h *HostApi) AssignHost(c *gin.Context) {
	response := NewResponse(c)

	var ha hostAssignment
	if err := c.ShouldBindJSON(&ha); err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
	}
	userId := ha.UserId
	// hostIds := ha.HostIds
	hostId := ha.HostId

	err := model.HostAssignment(userId, hostId)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
	}
	response.ToResponse(SuccessResponse{Msg: "分配主机成功"})
}

// SshHost
// @Tags HostApi
// @Summary SSH连接到主机
// @Description redirect to SSH access path for a host
// @Security Bearer
// @Param id path int true "Host ID"
// @Success 302 {string} string "Redirect to SSH access path"
// @Failure 404 {object} gin.H
// @Router /api/v1/hosts/{id}/ssh [get]
func (h *HostApi) SshHost(c *gin.Context) {
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
	sshAccessPath := "http://" + host.HostIP + ":3003/wetty"
	c.Redirect(http.StatusFound, sshAccessPath)
}

// https://github.com/yohanapriyandi/Golang-Rest-API-With-GIN-GORM.git
