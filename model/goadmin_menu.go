package model

type Menu struct {
	*Model
	Name       string `gorm:"column:name;NOT NULL" json:"name"`           // 菜单名称
	Icon       string `gorm:"column:icon" json:"icon"`                    // 图标
	Url        string `gorm:"column:url" json:"url"`                      // URL地址
	Pid        int    `gorm:"column:pid;default:0;NOT NULL" json:"pid"`   // 上级ID
	Type       int    `gorm:"column:type;default:0;NOT NULL" json:"type"` // 类型：1模块 2导航 3菜单 4节点
	IsShow     int    `gorm:"column:is_show;default:1" json:"is_show"`    // 是否显示：1显示 2不显示
	Sort       int    `gorm:"column:sort" json:"sort"`                    // 显示顺序
	Permission string `gorm:"column:permission" json:"permission"`        // 权限标识
	Remark     string `gorm:"column:remark" json:"remark"`                // 菜单备注
	Status     int    `gorm:"column:status;default:1" json:"status"`      // 状态1-在用 2-禁用
}

func (m *Menu) TableName() string {
	return "go_admin_" + "menu"
}
