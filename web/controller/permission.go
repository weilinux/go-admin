package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
)

func GetPermissions(c *gin.Context) {
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
