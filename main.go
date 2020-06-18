package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangyouwei/consulhandelnginx/configure"
	"github.com/yangyouwei/consulhandelnginx/gencnf"
	"github.com/yangyouwei/consulhandelnginx/jsonParse"
	"io/ioutil"
	"net/http"
)

func main()  {

	fmt.Println(configure.Mainconf)

	router := gin.Default()
	router.POST("/innotify", RequestHandler)
	router.Run(":9527") // listen and serve on 0.0.0.0:8080
}

func RequestHandler(c *gin.Context) {
	//read json
	res,err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "request error",
		})
	}
	//parse json get service name ,addr and port
	service,err := jsonParse.GetRes(string(res))
	if service.Servers == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "json error",
		})
	}
	//生成配置文件
	gencnf.GenConf(service)
	c.String(http.StatusOK, "ok")
}