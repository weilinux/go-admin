package app

import (
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"os"
)

// allowed app env name
const (
	EnvProd = "prod"
	EnvPre  = "pre"
	EnvTest = "test"
	EnvDev  = "dev"
)

const (
	DateFormat = "2006-01-02 15:04:05"

	configSuffix = ".toml"
)

var (
	Name    = "github.com/weilinux/go-gin-skeleton-auth"
	EnvName = EnvDev

	Debug    bool
	Hostname string

	GitInfo  model.AppInfo
	HttpPort = 9550
)

// the app work dir path
var WorkDir, _ = os.Getwd()

// IsEnv current env name check
func IsEnv(env string) bool {
	return env == EnvName
}
