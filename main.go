package main

import (
	"flag"
	"github.com/yangyouwei/consulhandelnginx/conflib"
	"github.com/yangyouwei/consulhandelnginx/loglib"
	"github.com/yangyouwei/consulhandelnginx/router"
)

func init()  {
	s := flag.String("c", "./conf.ini", "-c /etc/main.conf")
	flag.Parse()
	//解析参数
	if *s == "" {
		flag.Usage()
		panic("process exit!")
	}
	conflib.InitConf(s)
	loglib.InitLog()
}

//defaut listen 8000
//test url  http://127.0.0.1:8000/api/dosmoething0

func main()  {
	router.InitRouter()
}
