package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContApi struct {
	BaseApi
}

type contAssignment struct {
	UserId int64 `json:"user_id" binding:"required"`
	// ContIds []int64 `json:"cont_id" binding:"required"`
	ContId int64 `json:"cont_id" binding:"required"`
}

type Cont struct {
	ID     string   `json:"Id"`
	Names  []string `json:"Names"`
	Image  string   `json:"Image"`
	State  string   `json:"State"`
	Status string   `json:"Status"`
}

// Container 表示 Docker 容器信息
type Container struct {
	ID              string          `json:"Id"`
	Names           []string        `json:"Names"`
	Image           string          `json:"Image"`
	ImageID         string          `json:"ImageID"`
	Command         string          `json:"Command"`
	Created         int64           `json:"Created"`
	Ports           []Port          `json:"Ports"`
	Labels          Labels          `json:"Labels"`
	State           string          `json:"State"`
	Status          string          `json:"Status"`
	HostConfig      HostConfig      `json:"HostConfig"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
	Mounts          []Mount         `json:"Mounts"`
}

// Port 表示容器暴露的端口
type Port struct {
	IP          string `json:"IP,omitempty"`
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort,omitempty"`
	Type        string `json:"Type"`
}

// Labels 表示容器标签
type Labels map[string]string

// HostConfig 表示主机配置
type HostConfig struct {
	NetworkMode string `json:"NetworkMode"`
}

// NetworkSettings 表示网络设置
type NetworkSettings struct {
	Networks map[string]Network `json:"Networks"`
}

// Network 表示网络配置
type Network struct {
	IPAMConfig          interface{} `json:"IPAMConfig"`
	Links               interface{} `json:"Links"`
	Aliases             interface{} `json:"Aliases"`
	MacAddress          string      `json:"MacAddress"`
	NetworkID           string      `json:"NetworkID"`
	EndpointID          string      `json:"EndpointID"`
	Gateway             string      `json:"Gateway"`
	IPAddress           string      `json:"IPAddress"`
	IPPrefixLen         int         `json:"IPPrefixLen"`
	IPv6Gateway         string      `json:"IPv6Gateway"`
	GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
	DriverOpts          interface{} `json:"DriverOpts"`
	DNSNames            interface{} `json:"DNSNames"`
}

// Mount 表示容器挂载
type Mount struct {
	Type        string `json:"Type"`
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
	Propagation string `json:"Propagation"`
	Name        string `json:"Name,omitempty"`
	Driver      string `json:"Driver,omitempty"`
}

// GetConts
// @Tags ContApi
// @Summary 获取容器列表
// @Description get container list from Docker
// @Security Bearer
// @Success 200 {object} controller.SuccessResponse{data=map[string]interface{}}
// @Router /api/v1/conts [get]
func (h *ContApi) GetConts(c *gin.Context) {
	// TODO::接口响应慢的优化,已经记录文档
	response := NewResponse(c)

	// page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// rows, count := model.GetConts(page, limit)
	resp, err := http.Get("http://47.92.27.135:8088/containers/json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var containers []Container
	var conts []Cont
	if err := json.NewDecoder(resp.Body).Decode(&containers); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, container := range containers {
		// 创建一个新的 Cont 实例
		newCont := Cont{
			ID:     container.ID,
			Names:  container.Names,
			Image:  container.Image,
			State:  container.State,
			Status: container.Status,
		}
		conts = append(conts, newCont)
	}

	data := map[string]interface{}{
		"status": "UP",
		"rows":   conts,
		"total":  len(containers),
	}
	response.ToResponse(SuccessResponse{
		Data: data,
	})
}
