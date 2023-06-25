package model

import (
	"time"
)

// AppInfo app git info
//
//	{
//		"version": "@pkg_version",
//		"tag": "@pkg_branch_alias_version",
//		"releaseAt": "@pkg_release_date"
//	}
type AppInfo struct {
	Tag       string `json:"tag" description:"get tag name"`
	Version   string `json:"version" description:"git repo version."`
	ReleaseAt string `json:"releaseAt" description:"latest commit date"`
}

// 基底建立不好
type Model struct {
	ID          int       `gorm:"primary_key" json:"id"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time"`
}

// func (m *Model) BeforeUpdate(db *gorm.DB) error {
// 	n := &time.Time{}
// 	if err := n.Scan(time.Now()); err != nil {
// 		return err
// 	}
// 	db.Statement.SetColumn("updated_time", *n)
//
// 	return nil
// }
//
// func (m *Model) BeforeCreate(db *gorm.DB) error {
// 	n := &time.Time{}
// 	if err := n.Scan(time.Now()); err != nil {
// 		return err
// 	}
// 	db.Statement.SetColumn("created_time", *n)
// 	db.Statement.SetColumn("updated_time", *n)
//
// 	return nil
// }
//
// func (t *time.Time) Scan(v interface{}) error {
// 	switch vt := v.(type) {
// 	case time.Time:
// 		// 字符串转成 time.Time 类型
// 		*t = time.Time{}(vt)
// 	default:
// 		return errors.New("类型处理错误")
// 	}
// 	return nil
// }
