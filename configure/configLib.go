package configure

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
)

type MainConf struct {
	TplPath             string
	UpsPath          string
	Log bool
	LogPath string
}

var  mainconfigpath string = "./conf"
var Mainconf MainConf

func init()  {
	Mainconf.GetMainConf()
	fmt.Println(Mainconf)
}

func (m *MainConf) GetMainConf() {
	cfg, err := goconfig.LoadConfigFile(mainconfigpath)
	if err != nil {
		log.Println("读取配置文件失败[main.conf]")
		return
	}
	key, err := cfg.GetSection("main")
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
		}
	}
}

func stringtobool(s string) bool {
	if s == "true" {
		return true
	}
	return false
}