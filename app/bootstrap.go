package app

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/i18n"
	"github.com/weilinux/go-gin-skeleton-auth/model"
	"io/ioutil"
	"os"
	"strings"
)

var (
	I18n   *i18n.I18n
	Config *config.Config
)

func BootStrap(configDir string) {
	initAppInfo()
	loadAppConfig(configDir)
	initLogger()

	initLanguage()
}

func initAppInfo() {
	// git repo info
	GitInfo = model.AppInfo{}
	infoFile := "static/app.json"

	if fsutil.IsFile(infoFile) {
		err := jsonutil.ReadFile(infoFile, &GitInfo)
		if err != nil {
			color.Error.Println(err.Error())
		}
	}
}

func loadAppConfig(configDir string) {
	fmt.Printf("- work dir: %s\n", WorkDir)
	fmt.Printf("- config dir: %s\n", configDir)
	// fmt.Printf("- load config: conf/app.ini, %s\n", envFile)

	files, err := getConfigFiles(configDir)
	if err != nil {
		color.Error.Println(err.Error())
		os.Exit(1)
	}

	// config instance
	Config = config.Default()
	Config.AddDriver(toml.Driver)

	err = Config.LoadFiles(files...)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Name = Config.String("name")
	Debug = Config.Bool("debug")

}

// 获取某个文件夹下的配置文件列表
func getConfigFiles(confDir string) ([]string, error) {
	var files = make([]string, 0)

	fileInfoList, err := ioutil.ReadDir(confDir)
	if err != nil {
		return files, err
	}

	pathSep := string(os.PathSeparator)
	// app.toml is must exist
	baseFile := confDir + pathSep + "app" + configSuffix
	files = append(files, baseFile)

	// _dev.toml
	suffix := "-" + EnvName + configSuffix
	for _, f := range fileInfoList {
		// app_dev.toml
		if !f.IsDir() && strings.HasSuffix(f.Name(), suffix) {
			files = append(files, confDir+pathSep+f.Name())
		}
	}

	return files, err
}

// func loadAppTemplate() {
// 	Tpl = template.Must(template.ParseGlob("templates/*"))
// }
