package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Role struct {
	*Model
	Name   string `gorm:"column:name;NOT NULL" json:"name"`      // 角色名称
	Code   string `gorm:"column:code;NOT NULL" json:"code"`      // 角色编码
	Sort   int    `gorm:"column:sort;default:0" json:"sort"`     // 排序
	Status int    `gorm:"column:status;default:1" json:"status"` // 状态 1-启用 2-禁用
}

func (r *Role) TableName() string {
	return "go_admin_" + "role"
}

// TODO: db.WithContext(ctx)
func FindRoleByName(roleName string) (Role, error) {
	var role Role
	result := db.Where("name = ?", roleName).First(&role)
	if result.RowsAffected < 1 {
		return role, errors.New("role does not exist")
	}
	return role, nil
}

func CreateRole(role Role) (Role, error) {
	result := db.Create(&role)
	if result.Error != nil {
		return Role{}, result.Error
	}
	return role, result.Error
}

// List 列表
func GetAllRoles() []Role {
	var roles []Role
	db.Find(&roles)
	return roles
}

func GetRoleById(Id int64) (*Role, *gorm.DB) {
	var role Role
	db := db.Where("id = ?", Id).Find(&role)
	return &role, db
}

func DeleteRole(Id int64) (Role, error) {
	var getRole Role
	// 物理删除 ,gorm软删除，硬删除, Unscoped无关，这里必须是指针
	if err := db.Where("id = ?", Id).Delete(&getRole).Error; err != nil {
		return Role{}, err
	}
	return getRole, nil
}
