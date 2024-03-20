package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"strconv"
	"time"
)

func DeleteRole(c *gin.Context) {
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

func EditRole(c *gin.Context) {
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

func AddRole(c *gin.Context) {
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

func RoleInfo(c *gin.Context) {
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

func GetRoles(c *gin.Context) {
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
// func UpdateRole(c *fiber.Ctx) error {
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
