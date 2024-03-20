package model

type Permission struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (p *Permission) TableName() string {
	return "go_admin_" + "permission"

}

func AllPermissions() []Permission {
	var permissions []Permission
	db.Find(&permissions)
	return permissions
}
