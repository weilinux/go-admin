package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

//	func MonitorHosts(c *gin.Context) {
//		id := c.Param("id")
//		ID, err := strconv.ParseInt(id, 10, 64)
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{
//				"error": err.Error(),
//			})
//		}
//		hostItem, _ := model.GetHostById(ID)
//
//		// fetch hostInfo metrics
//		hostInfo, err := host.Info()
//		if err != nil {
//			fmt.Println("get host info fail, error: ", err)
//		}
//		data := map[string]interface{}{
//			"host":     hostItem,
//			"hostInfo": hostInfo,
//		}
//
//		c.JSON(http.StatusOK, data)
//	}

// type InfoStat struct {
func MonitorHosts(c *gin.Context) {
	response := NewResponse(c)
	resp, err := http.Get("http://47.92.27.135:3000/api/metrics")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var cpu map[string]float32

	err = json.Unmarshal(body, &cpu)
	if err != nil {
		fmt.Println(err.Error())
	}

	response.ToResponse(SuccessResponse{
		Msg:  "135",
		Data: cpu,
	})
}

// type InfoStat struct {
//     Hostname             string `json:"hostname"`
//     Uptime               uint64 `json:"uptime"`
//     BootTime             uint64 `json:"bootTime"`
//     Procs                uint64 `json:"procs"`           // number of processes
//     OS                   string `json:"os"`              // ex: freebsd, linux
//     Platform             string `json:"platform"`        // ex: ubuntu, linuxmint
//     PlatformFamily       string `json:"platformFamily"`  // ex: debian, rhel
//     PlatformVersion      string `json:"platformVersion"` // version of the complete OS
//     KernelVersion        string `json:"kernelVersion"`   // version of the OS kernel (if available)
//     KernelArch           string `json:"kernelArch"`      // native cpu architecture queried at runtime, as returned by `uname -m` or empty string in case of error
//     VirtualizationSystem string `json:"virtualizationSystem"`
//     VirtualizationRole   string `json:"virtualizationRole"` // guest or host
//     HostID               string `json:"hostid"`             // ex: uuid
// }

func SystemResources() map[string]interface{} {
	result := map[string]interface{}{}
	percent, _ := cpu.Percent(time.Second*3, true)
	result["cpu_used_percent"] = percent
	memInfo, _ := mem.VirtualMemory()
	result["memory_used_percent"] = memInfo.UsedPercent

	// 磁盘
	// disks := []map[string]interface{}{}
	var disks []map[string]interface{}
	parts, _ := disk.Partitions(true)
	for _, part := range parts {
		if strings.HasPrefix(part.Device, "/dev/") {
			usage, _ := disk.Usage(part.Mountpoint)
			disks = append(disks, map[string]interface{}{"device": part.Device, "mount_point": part.Mountpoint, "used_percent": usage.UsedPercent, "inodes_used_percent": usage.InodesUsedPercent})
		}
	}
	result["disks"] = disks

	// 网络
	interfaces, _ := net.IOCounters(true)
	// nets := []map[string]interface{}{}
	var nets []map[string]interface{}
	for _, inte := range interfaces {
		nets = append(nets, map[string]interface{}{"name": inte.Name, "bytes_sent": inte.BytesSent, "bytes_recv": inte.BytesRecv})
	}
	result["net"] = nets
	return result

}

// result data struct
// {
// "cpu_used_percent": [
// 7.046979865786559,
// 4.304635761596584,
// 2.01342281884848
// ],
// "disks": [
// {
// "device": "/dev/sda1",
// "inodes_used_percent": 6.301271950879583,
// "mount_point": "/",
// "used_percent": 19.380184817204057
// }
// ],
// "memory_used_percent": 41.11186087644503,
// "net": [
// {
// "bytes_recv": 6035882499,
// "bytes_sent": 6035882499,
// "name": "lo"
// },
// {
// "bytes_recv": 897458846,
// "bytes_sent": 5788731489,
// "name": "eth0"
// },
// {
// "bytes_recv": 404561,
// "bytes_sent": 69195902,
// "name": "docker0"
// }
// ]
// }
