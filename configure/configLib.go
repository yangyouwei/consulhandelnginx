package configure

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"path/filepath"
)

type MainConf struct {
	TplPath             string
	UpsPath          string
	Log bool
	LogPath string
	DyupsUrl string
	ListenPort string
}

var Mainconf MainConf

func Initconfig(s *string)  {
	cstr, err := filepath.Abs(*s)
	if err != nil {
		log.Panicln(err)
	}
	//初始化配置文件
	cfg, err := goconfig.LoadConfigFile(cstr)
	if err != nil {
		log.Println("读取配置文件失败[config.ini]")
		log.Panic(err)
	}
	Mainconf.GetMainConf(cfg)
}

func (m *MainConf) GetMainConf(cfg *goconfig.ConfigFile) {
	key, err := cfg.GetSection("main")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(key)
	for k, v := range key {
		switch k {
		case "tpl_path":
			m.TplPath = v
		case "upstream_path":
			m.UpsPath = v
		case "log":
			m.Log = stringtobool(v)
		case "log_dir":
			m.LogPath = v
		case "dyups_url":
			m.DyupsUrl = v
		case "listen_port":
			m.ListenPort = v
		}
	}
}

func stringtobool(s string) bool {
	if s == "true" {
		return true
	}
	return false
}