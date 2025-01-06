package request

type EditHost struct {
	HostName     string `form:"HostName"`
	HostPort     int    `form:"HostPort"`
	HostUser     string `form:"HostUser"`
	HostPassword string ``
	HostIP       string `form:"HostIP"`
}
