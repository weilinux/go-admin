package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/weilinux/go-gin-skeleton-auth/model/mysql"
)

// Apart from being an excellent ORM for Go developers, its built to be developer-friendly
// and easy to comprehend. Gorm is built on top of the database/sql packages.
// overview and features of the ORM are:
// Developer Friendly
// Every feature comes with a tests
// Hooks / Callbacks (Before/After Create/Save/Update/Delete/Find)
// Eager loading with Preload, Joins
// Context, Prepared Statement Mode, DryRun Mode
// SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
// Transactions, Nested Transactions, Save Point, Rollback To to Saved Point
// Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism)
// SQL Builder
// Logger

type User struct {
	gorm.Model
	UserName string
	Password string
}

func init() {
	mysql.Connect()
	DBM = mysql.GetDB()
	DBM.AutoMigrate(&User{})
}

func FindUserByName(userName string) (User, error) {
	var user User
	result := DBM.Where("user_name = ?", userName).First(&user)
	if result.RowsAffected < 1 {
		return user, errors.New("user does not exist")
	}
	return user, nil
}

func CreateUser(user User) (User, error) {
	// TODO: 使用GORM的方式来创建记录，参考链接https://medium.com/@itskenzylimon/getting-started-on-golang-gorm-af49381caf3f
	// TODO: 使用GORM的方式来创建记录，参考链接https://medium.com/@itskenzylimon/getting-started-on-golang-gorm-af49381caf3f
	result := DBM.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, result.Error
}

func GetAllUsers() []User {
	var users []User
	DBM.Find(&users)
	return users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	db := DBM.Where("id = ?", Id).Find(&user)
	return &user, db
}

func DeleteUser(Id int64) User {
	var getUser User
	if err := DBM.Where("id = ?", Id).Delete(getUser).Error; err != nil {
		fmt.Println(err.Error())
	}
	return getUser
}

// TODO: add support https://gitee.com/GoProgect/ginBlog/blob/master/model/User.go
// TODO: add support https://gitee.com/GoProgect/ginBlog/blob/master/model/User.go
// TODO: add support https://gitee.com/GoProgect/ginBlog/blob/master/model/User.go

// func UpdateUser(username string, user User) User {
//
// }

func UserExists(userInput User) bool {
	if _, err := FindUserByName(userInput.UserName); err != nil {
		return false
	}
	return true
}
