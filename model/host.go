package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/weilinux/go-gin-skeleton-auth/model/mysql"
	"net"
)

// Host 代码的形式化越优秀，说明代码越符合代码级别的CURD

var DBM *gorm.DB

var DbM *sql.DB

type Host struct {
	gorm.Model
	HostName     string
	HostPort     int
	HostUser     string
	HostPassword string
	HostIP       string
	isAlive      bool
	Users        []User `gorm:"many2many:go_admin_host_user"`
}

func (h *Host) TableName() string {
	return "go_admin_" + "host"
}

func init() {
	mysql.Connect()
	DBM = mysql.GetDB()
	DBM.AutoMigrate(&Host{})
	DBM.AutoMigrate(&User{})
	DbM = mysql.GetDb()
}

func (h *Host) IsAlive() bool {
	// send alert msg to Email 421596839@qq.com
	// 超时则显示状态为unknown
	target := fmt.Sprintf("%s:%d", h.HostIP, h.HostPort)
	conn, err := net.Dial("tcp", target)
	if err != nil {
		h.isAlive = false
		return false
	}
	if conn != nil {
		_ = conn.Close()
		h.isAlive = true
		return true
	}
	h.isAlive = false
	return false
}

func (h Host) Connect() error {
	return nil
}

func (h Host) audit() error {
	// crond running
	fmt.Println(h)
	return nil
}

// CreateHost CURD
func (h *Host) CreateHost() *Host {
	if DBM.NewRecord(h) {
		if count := DBM.Create(&h).RowsAffected; count != 1 {
			fmt.Printf("new record failed\n")
			return nil
		}
	} else {
		return nil
	}
	return h
}

// GetUserHosts CURD ,
func GetUserHosts(id int, page, limit int) ([]Host, int) {
	// https://github.com/go-gorm/gorm/issues/2994
	var hosts []Host
	var count int

	if err := DBM.
		Table("go_admin_host").
		Joins(" INNER JOIN go_admin_host_user on go_admin_host.id = go_admin_host_user.host_id ").
		Where("go_admin_host_user.user_id = ? AND go_admin_host.deleted_at is NULL", id).
		Count(&count).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&hosts).Error; err != nil {
		fmt.Printf("Query: %v\n", err)
	}

	return hosts, count
}

// GetUserUnbindHosts 如果没有用户id的话来进行全局查找, 所以这里不需要这个新的函数
func GetUserUnbindHosts(id int64) []Host {
	var hosts []Host

	sub := DBM.Table("go_admin_host_user").Select("host_id").Where("user_id = ?", id).SubQuery()
	if err := DBM.
		Table("go_admin_host").
		Where("deleted_at is NUll AND id NOT IN ?", sub).
		Find(&hosts).Error; err != nil {
		fmt.Printf("Query: %v\n", err)
	}
	return hosts
}

func GetHostById(Id int64) (*Host, *gorm.DB) {
	var host Host
	// 这个是系统的db的时候，为什么会导致记录更新之后只有一条了, 要重要定义一个db,不是全局的db
	db := DBM.Where("id = ?", Id).Find(&host)
	return &host, db
}

// DeleteHost CURD
func DeleteHost(Id int64) Host {
	var getHost Host
	if err := DBM.Where("id = ?", Id).Delete(getHost).Error; err != nil {
		fmt.Println(err.Error())
	}
	return getHost
}

func HostAssignment(userId int64, hostIds []int64) error {
	sqlStr := `insert into go_admin_host_user(host_id, user_id) values(?, ?)`
	for _, hostId := range hostIds {
		_, err := mysql.Db.Exec(sqlStr, hostId, userId)
		if err != nil {
			fmt.Printf("Insert failed: %v", err)
			return err
		}

	}
	return nil
}
