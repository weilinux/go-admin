package model

import "gorm.io/gorm"

// 分页基底
type Entity interface {
	Count(db *gorm.DB) int64
	Take(db *gorm.DB, limit int, offset int) interface{}
}
