package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"strconv"
	"time"
)

type RoleApi struct {
	BaseApi
}

// DeleteRole
// @Tags RoleApi
// @Summary 删除角色
// @Description 根据角色ID删除角色
// @Security Bearer
// @Param id path int true "角色ID"
// @Success 200 {object} controller.SuccessResponse{msg=string}
// @Router /api/v1/roles/{id} [delete]
func (r *RoleApi) DeleteRole(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	if _, err = model.DeleteRole(ID); err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
	}
	response.ToResponse(SuccessResponse{
		Msg: "删除角色成功",
	})
}

// EditRole
// @Tags RoleApi
// @Summary 编辑角色
// @Description 根据角色ID编辑角色信息
// @Security Bearer
// @Param id path int true "角色ID"
// @Param RoleName formData string false "角色名称"
// @Param Password formData string false "角色密码"
// @Success 200 {object} controller.SuccessResponse{msg=string, data=map[string]interface{}}
// @Router /api/v1/roles/{id} [put]
func (r *RoleApi) EditRole(c *gin.Context) {
	// response := NewResponse(c)
	// ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	// role := &model.Role{Model: &model.Model{ID: int(ID)}}

	// role.RoleName = c.PostForm("RoleName")
	// role.Password = c.PostForm("Password")

	// roleDetails, _ := model.GetRole(ID)
	// if role.RoleName != "" {
	// 	roleDetails.RoleName = role.RoleName
	// }
	// if role.Password != "" {
	// 	roleDetails.Password = role.Password
	// }

	// TODO: 修改角色成功后, update time没有更新!

	// db.Save(&roleDetails)
	// model.SaveRole(roleDetails)

	// response.ToResponse(SuccessResponse{
	// 	Msg:  "修改角色配置成功",
	// 	Data: roleDetails,
	// })
}

// AddRole
// @Tags RoleApi
// @Summary 添加角色
// @Description 添加一个新的角色
// @Security Bearer
// @Param name formData string true "角色名称"
// @Param status formData int true "角色状态"
// @Param code formData string true "角色代码"
// @Param sort formData int true "排序值"
// @Success 200 {object} controller.SuccessResponse{msg=string, data=map[string]interface{}}
// @Router /api/v1/roles [post]
func (r *RoleApi) AddRole(c *gin.Context) {
	response := NewResponse(c)
	var role = &model.Role{}
	role.Name = c.PostForm("name")
	role.Status, _ = strconv.Atoi(c.PostForm("status"))
	role.Code = c.PostForm("code")
	role.Sort, _ = strconv.Atoi(c.PostForm("sort"))

	_, err := model.CreateRole(model.Role{
		Model: &model.Model{
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		},
		Name:   role.Name,
		Status: role.Status,
		Code:   role.Code,
		Sort:   role.Sort,
	})
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(SuccessResponse{
		Msg:  "添加角色成功",
		Data: role,
	})
}

// RoleInfo
// @Tags RoleApi
// @Summary 获取角色详情
// @Description 根据角色ID获取角色的详细信息
// @Security Bearer
// @Param id path int true "角色ID"
// @Success 200 {object} controller.SuccessResponse{msg=string, data=map[string]interface{}}
// @Router /api/v1/roles/{id} [get]
func (r *RoleApi) RoleInfo(c *gin.Context) {
	response := NewResponse(c)
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.ToErrorResponse(errcode.NotFound.WithDetails(err.Error()))
		return
	}

	role := model.Role{
		Model: &model.Model{
			ID: uint(ID),
		},
	}

	count := model.GetRole(&role)
	if count < 1 {
		response.ToErrorResponse(errcode.NotFound.WithDetails("角色不存在"))
		return
	}
	response.ToResponse(SuccessResponse{Data: role, Msg: "角色详情"})
}

// GetRoles
// @Tags RoleApi
// @Summary 获取角色列表
// @Description 获取所有角色的列表
// @Security Bearer
// @Success 200 {object} controller.SuccessResponse{msg=string, data=map[string]interface{}}
// @Router /api/v1/roles [get]
func (r *RoleApi) GetRoles(c *gin.Context) {
	response := NewResponse(c)
	roles := model.AllRoles()
	data := map[string]interface{}{
		"message": "角色管理界面",
		"rows":    roles,
	}

	response.ToResponse(SuccessResponse{
		Msg:  "角色管理界面",
		Data: data,
		Code: 200,
	})
}

// 注意查看这个更新用户的代码
// func  (r *RoleApi) UpdateRole(c *fiber.Ctx) error {
// 	if err := middlewares.IsAuthorize(c, "roles"); err != nil {
// 		return err
// 	}
//
// 	id, _ := strconv.Atoi(c.Params("id"))
//
// 	var roleDto fiber.Map
//
// 	if err := c.BodyParser(&roleDto); err != nil {
// 		return err
// 	}
//
// 	list := roleDto["permissions"].([]interface{})
//
// 	permissions := make([]models.Permission, len(list))
//
// 	for i, permissionId := range list {
// 		id, _ := strconv.Atoi(permissionId.(string))
// 		permissions[i] = models.Permission{
// 			Id: uint(id),
// 		}
// 	}
//
// 	var result interface{}
// 	database.DB.Table("role_permissions").Where("role_id", id).Delete(result)
//
// 	role := models.Role{
// 		Id:          uint(id),
// 		Name:        roleDto["name"].(string),
// 		Permissions: permissions,
// 	}
//
// 	database.DB.Model(&role).Updates(role)
//
// 	return c.JSON(role)
// }
