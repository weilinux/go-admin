package mysql

import (
	"fmt"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	// db *sql.DB
	db *gorm.DB
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:111000!Wei@tcp(47.92.27.135:3306)/godb2?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败!")
	}

	tables := model.Initialize(conn)
	SyncTables(conn, tables)
	sqlDB, err := conn.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	return conn
}

// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
func SyncTables(db *gorm.DB, tables []interface{}) {
	if err := db.AutoMigrate(tables...); err != nil {
		fmt.Println(err.Error())
		panic("初始化表结构失败!")
	}

	// delete fields
	// db.Migrator().DropColumn(&model.Menu{}, "deleted_time")
	// db.Migrator().DropColumn(&model.User{}, "deleted_time")
	// db.Migrator().DropColumn(&model.Host{}, "deleted_time")
	// db.Migrator().DropColumn(&model.Role{}, "deleted_time")
}

func CloseDatabase(conn *gorm.DB) error {
	sqlDB, _ := conn.DB()
	return sqlDB.Close()
}
