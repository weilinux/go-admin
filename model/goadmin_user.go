package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// Apart from being an excellent ORM for:uu Go developers, it's built to be developer-friendly
// and easy to comprehenkkkd. Gorm is built on top of the database/sql packages.
// overview and features of the ORM are:
// Developer Friendly
// Every feature comes with a tests
// Hooks / Callbacks (Before/After Create/Save/Update/Delete/Find)
// Eager loading with Preload, Joins
// Context, Prepared Statement Mode, DryRun Mode
// SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
// Transactions, Nested Transactions, Save Point, Rollback to Saved Point
// Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism)
// SQL Builder
// Logger

type User struct {
	*Model
	UserName string
	Password string `json:"-"`
	RoleID   uint
	Role     Role `json:"role"`
	// RoleID uint `gorm:"column:role_id;" json:"role_id"`
	// Role   Role `gorm:"foreignKey:RoleID;References:Id" json:"role"`
}

func (u *User) TableName() string {
	return "go_admin_" + "user"
}

func (u *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)
	return total
}

func (u *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User
	db.Table("go_admin_user").Preload("Role").Offset(offset).Limit(limit).Find(&users)
	return users
}

func FindUserByName(userName string) (User, error) {
	var user User
	result := db.Where("user_name = ?", userName).First(&user)
	if result.RowsAffected < 1 {
		return user, errors.New("user does not exist")
	}
	return user, nil
}

func CreateUser(user User) (User, error) {
	// TODO: 使用GORM的方式来创建记录，参考链接https://medium.com/@itskenzylimon/getting-started-on-golang-gorm-af49381caf3f
	// TODO: 使用GORM的方式来创建记录，参考链接https://medium.com/@itskenzylimon/getting-started-on-golang-gorm-af49381caf3f
	result := db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, result.Error
}

// Optimi: 优化用户查询
// user := models.User{
// Id: uint(id),
// }
// database.DB.Preload("Role").Find(&user)

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	db := db.Preload("Role").Where("id = ?", Id).Find(&user)
	return &user, db
}

func DeleteUser(Id int64) User {
	var getUser User
	// 物理删除 ,gorm软删除，硬删除, Unscoped无关，这里必须是指针
	if err := db.Where("id = ?", Id).Delete(&getUser).Error; err != nil {
		fmt.Println(err.Error())
	}
	return getUser
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {
	// // var user User
	// var maps = make(map[string]interface{})
	// maps["password"] = data.Password
	//
	// if err := db.Select("password").Where("id = ?", id).Updates(&data).Error; err != nil {
	// 	return errmsg.ERROR
	// }
	// return errmsg.SUCCSE
	return 0
}

// TODO: add support https://gitee.com/GoProgect/ginBlog/blob/master/model/User.go
// TODO: add support https://gitee.com/GoProgect/ginBlog/blob/master/model/User.go
// TODO: add support https://gitee.com/GoProgect/ginBlog/blob/master/model/User.go

func UserExists(userInput User) bool {
	if _, err := FindUserByName(userInput.UserName); err != nil {
		return false
	}
	return true
}
