package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
)

type PermissionApi struct {
	BaseApi
}

// GetPermissions
// @Tags PermissionApi
// @Summary 获取权限列表
// @Description 获取所有权限的列表
// @Security Bearer
// @Success 200 {object} controller.SuccessResponse{data=map[string]interface{}}
// @Router /api/v1/permissions [get]
func (p *PermissionApi) GetPermissions(c *gin.Context) {
	response := NewResponse(c)
	permission := model.AllPermissions()
	data := map[string]interface{}{
		"message": "权限管理界面",
		"rows":    permission,
	}

	response.ToResponse(SuccessResponse{
		Msg:  "权限管理界面",
		Data: data,
		Code: 200,
	})
}
