package model

type UserRole struct {
	*Model
	UserId int `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"` // 用户id
	RoleId int `gorm:"column:role_id;default:0;NOT NULL" json:"role_id"` // 角色id
}

func (ur *UserRole) TableName() string {
	return "go_admin_" + "user_role"
}

// UserRole CURD ops
