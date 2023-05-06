package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"net/http"
)

// UserManage 好的代码， 你是在添加新代码，而不是修改原有代码。添加新代码，却不是流水账式代码
// UserManage 好的代码， 你是在添加新代码，而不是修改原有代码。添加新代码，却不是流水账式代码
// UserManage 好的代码， 你是在添加新代码，而不是修改原有代码。添加新代码，却不是流水账式代码

// GetUsers TODO: add 怎么做用户列表分布呢，使用了两个方法获取，参考ginBlog中查询用户列表函数 GetUsers()
// GetUsers TODO: add 怎么做用户列表分布呢，使用了两个方法获取，参考ginBlog中查询用户列表函数 GetUsers()
// GetUsers TODO: add 怎么做用户列表分布呢，使用了两个方法获取，参考ginBlog中查询用户列表函数 GetUsers()

func GetUsers(c *gin.Context) {
	users := model.GetAllUsers()
	data := map[string]interface{}{
		"message": "用户管理界面",
		"rows":    users,
	}
	c.JSON(http.StatusOK, data)
}
