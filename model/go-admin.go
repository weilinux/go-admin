package model

// User2 https://github.com/mwqnice/oh-admin/blob/main/internal/model/admin_user.go
// 不需要下面的结构体定义了，为什么我需要执行oh-admin.sql才能生产表的呢?
type User2 struct {
	Id          int    `:"id"`           // 用户id
	Realname    string `:"realname"`     // 真实姓名
	Username    string `:"username"`     // 登录用户名
	Gender      int    `:"gender"`       // 性别:1男 2女 3保密
	Avatar      string `:"avatar"`       // 头像
	Mobile      string `:"mobile"`       // 手机号
	Email       string `:"email"`        // 邮箱地址
	Address     string `:"address"`      // 地址
	Password    string `:"password"`     // 登录密码
	Salt        string `:"salt"`         // 盐加密
	Intro       string `:"intro"`        // 备注
	Status      int    `:"status"`       // 状态：1正常 2禁用
	LoginNum    int    `:"login_num"`    // 登录次数
	LoginIp     string `:"login_ip"`     // 最近登录ip
	LoginTime   int    `:"login_time"`   // 最近登录时间
	CreatedTime string `:"created_time"` // 创建时间
	UpdatedTime string `:"updated_time"` // 修改时间
}

type Menu struct {
	Id          int    `:"id"`           // 用户id
	Name        string `:"name"`         // 菜单名称
	Icon        string `:"icon"`         // 图标
	Url         string `:"url"`          // URL地址
	Pid         int    `:"pid"`          // 上级ID
	Type        int    `:"type"`         // 类型：1模块 2导航 3菜单 4节点
	Permission  string `:"permission"`   // 权限标识
	IsShow      int    `:"is_show"`      // 是否显示：1显示 2不显示
	Sort        int    `:"sort"`         // 显示顺序
	Remark      string `:"remark"`       // 菜单备注
	Status      int    `:"status"`       // 状态1-在用 2-禁用
	CreatedTime string `:"created_time"` // 创建时间
	UpdatedTime string `:"updated_time"` // 修改时间
}

type RoleMenu struct {
	Id          int    `:"id"`           // 用户id
	MenuId      int    `:"menu_id"`      // 菜单id
	RoleId      int    `:"role_id"`      // 角色id
	CreatedTime string `:"created_time"` // 创建时间
	UpdatedTime string `:"updated_time"` // 修改时间
}
