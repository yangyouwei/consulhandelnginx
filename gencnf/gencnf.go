package gencnf

import (
	"bytes"
	"fmt"
	"github.com/yangyouwei/consulhandelnginx/configure"
	"github.com/yangyouwei/consulhandelnginx/jsonParse"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)


func ReadTpl(p string) string {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}


func GenConf(s *jsonParse.Service)  {
	// 读取模板
	templateText := ReadTpl(configure.Mainconf.TplPath)
	buffer := new(bytes.Buffer)
	t := template.Must(template.New("upstream").Parse(templateText))
	err := t.Execute(buffer, s)
	if err != nil {
		log.Println("Executing template:", err)
	}

	//创建配置文件
	fd, err := os.OpenFile(configure.Mainconf.UpsPath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	fd.Write(buffer.Bytes())
	fmt.Println("save conf to file.")
	fmt.Println("services: ",s)
}