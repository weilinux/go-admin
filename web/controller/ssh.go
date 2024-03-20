package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/weilinux/go-gin-skeleton-auth/model/connections"
	"github.com/weilinux/go-gin-skeleton-auth/pkg/errcode"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ShellWs(c *gin.Context) {
	response := NewResponse(c)
	var err error
	msg := c.DefaultQuery("msg", "")
	cols := c.DefaultQuery("cols", "150")
	rows := c.DefaultQuery("rows", "35")
	col, _ := strconv.Atoi(cols)
	row, _ := strconv.Atoi(rows)
	terminal := connections.Terminal{
		Columns: uint32(col),
		Rows:    uint32(row),
	}
	sshClient, err := connections.DecodedMsgToSSHClient(msg)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	if sshClient.IpAddress == "" || sshClient.Password == "" {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails("IP地址或密码不能为空!"))
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.ToErrorResponse(errcode.Fail.WithDetails(err.Error()))
		return
	}
	err = sshClient.GenerateClient()
	if err != nil {
		_ = conn.WriteMessage(1, []byte(err.Error()))
		_ = conn.Close()
		return
	}
	sshClient.RequestTerminal(terminal)
	sshClient.Connect(conn)
}
