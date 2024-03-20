package model

type Role struct {
	*Model
	// *gorm.Model
	// Id     uint   `json:"id"`
	Name        string       `gorm:"column:name;NOT NULL" json:"name"`                      // 角色名称
	Code        string       `gorm:"column:code;NOT NULL" json:"code"`                      // 角色编码
	Sort        int          `gorm:"column:sort;default:0" json:"sort"`                     // 排序
	Status      int          `gorm:"column:status;default:1" json:"status"`                 // 状态 1-启用 2-禁用
	Permissions []Permission `gorm:"many2many:go_admin_role_permission;" json:"permission"` // 角色权限
	// Permissions []Permission `gorm:"many2many:go_admin_role_permission;foreignKey:ID;references:Id" json:"permission"` // 角色权限
	// Permissions []Permission `gorm:"many2many:go_admin_role_permission;foreignKey:ID;JoinForeignKey:role_id;references:Id;JoinReferences:permission_id"`
}

func (r *Role) TableName() string {
	return "go_admin_" + "role"
}

func AllRoles() []Role {
	var roles []Role
	db.Preload("Permissions").Find(&roles)
	return roles
}

func GetRole(role *Role) int64 {
	count := db.Preload("Permissions").Where("id = ?", role.ID).Find(&role).RowsAffected
	return count
}

func CreateRole(role Role) (Role, error) {
	result := db.Create(&role)
	if result.Error != nil {
		return Role{}, result.Error
	}
	return role, result.Error
}

func DeleteRole(Id int64) (Role, error) {
	var getRole Role
	// 物理删除 ,gorm软删除，硬删除, Unscoped无关，这里必须是指针
	if err := db.Where("id = ?", Id).Delete(&getRole).Error; err != nil {
		return Role{}, err
	}
	return getRole, nil
}

// 保存用户角色信息
func SaveRole(role *Role) error {
	return db.Save(role).Error
}

// 更新用户角色
