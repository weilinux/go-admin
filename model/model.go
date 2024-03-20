package model

import (
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

type AppInfo struct {
	Tag       string `json:"tag" description:"get tag name"`
	Version   string `json:"version" description:"git repo version."`
	ReleaseAt string `json:"releaseAt" description:"latest commit date"`
}

// ID int64  `gorm:"column:id;AUTO_INCREMENT;comment:主键" `
// ID        uint `gorm:"primary_key"` from Tencent

type Model struct {
	ID          uint      `gorm:"primarykey;" json:"id"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time"`
	// DeletedTime time.Time `json:"deleted_time" gorm:"column:deleted_time"`
}

// modelDesign -> migrate-> modelController->modelCURD
func Initialize(mysql *gorm.DB) []interface{} {
	db = mysql
	return []interface{}{
		&Menu{},
		&User{},
		&Role{},
		&Permission{},
		&Host{},
	}
	// gorm.Migrator.DropColumn(&Menu{}, "deleted_time")
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
