package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/weilinux/go-gin-skeleton-auth/model"
)

var (
	Db *sql.DB
	db *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:111000!Wei@tcp(47.92.27.135:3306)/godb2"
	// 去初始化全局的db对象而不是新声明一个db变量
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		fmt.Printf("connect to Db failed, err:%v\n", err)
		return
	}

	// Db.SetConnMaxLifetime(time.Second*10)
	Db.SetMaxOpenConns(100)
	Db.SetMaxIdleConns(10)
	return
}

func ConnectDatabase() {
	mysql, err := gorm.Open("mysql", "root:111000!Wei@tcp(47.92.27.135:3306)/godb2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败!")
	}

	tables := model.Initialize(mysql)
	SyncTables(mysql, tables)
	db = mysql
}

func SyncTables(db *gorm.DB, tables []interface{}) {
	if err := db.AutoMigrate(tables...).Error; err != nil {
		fmt.Println(err.Error())
		panic("初始化表结构失败!")
	}
}

func CloseDatabase() error {
	return db.DB().Close()
}
