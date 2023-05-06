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
	HostIP       string
	HostPort     int
	HostUser     string
	HostPassword string
	isAlive      bool
	Users        []User `gorm:"many2many:host_users"`
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

	// if h.IsAlive() {
	// 	return nil
	// }
	// if user have privileges ? {
	//    then conn, err := net.Dial("tcp", h.Host)
	//    return nil
	// }
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

// GetAllHosts CURD ,
func GetAllHosts(id uint) []Host {
	// TODO: support gorm query style
	var hosts []Host
	// DBM.Find(&hosts)
	sqlStr := `SELECT id, host_name, host_port, host_user, host_password, host_ip from hosts WHERE id in (SELECT host_id from host_users WHERE user_id = ? and deleted_at IS NULL);`
	rows, err := mysql.Db.Query(sqlStr, id)
	// rows, err := mysql.Db.Query(sqlStr, 1)
	if err != nil {
		fmt.Printf("Query: %v\n", err)
	}
	for rows.Next() {
		var host Host
		err = rows.Scan(&host.ID, &host.HostName, &host.HostPort, &host.HostUser, &host.HostPassword, &host.HostIP)
		if err != nil {
			fmt.Println("scan error:", err)
		}
		hosts = append(hosts, host)
	}
	return hosts
}

// GetAllHostAll 如果没有用户id的话来进行全局查找, 所以这里不需要这个新的函数
func GetAllHostAll(id int64) []Host {
	// TODO: support gorm query style
	var hosts []Host
	// DBM.Find(&hosts)
	// sqlStr := `SELECT id, host_name, host_port, host_user, host_password, host_ip from hosts WHERE deleted_at IS NULL;`
	sqlStr := `SELECT id, host_name, host_port, host_user, host_password, host_ip from hosts WHERE deleted_at IS NULL and id not in (SELECT host_id from host_users WHERE user_id = ?);`
	// rows, err := mysql.Db.Query(sqlStr, id)
	rows, err := mysql.Db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("Query: %v\n", err)
	}
	for rows.Next() {
		var host Host
		err = rows.Scan(&host.ID, &host.HostName, &host.HostPort, &host.HostUser, &host.HostPassword, &host.HostIP)
		if err != nil {
			fmt.Println("scan error:", err)
		}
		hosts = append(hosts, host)
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
	sqlStr := `insert into host_users(host_id, user_id) values(?, ?)`
	for _, hostId := range hostIds {
		_, err := mysql.Db.Exec(sqlStr, hostId, userId)
		if err != nil {
			fmt.Printf("Insert failed: %v", err)
			return err
		}

	}
	return nil
}
