package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangyouwei/consulhandelnginx/curldyups"
	"github.com/yangyouwei/consulhandelnginx/gencnf"
	"github.com/yangyouwei/consulhandelnginx/jsonParse"
	"io/ioutil"
	"net/http"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Any("/innotify", RequestHandler)
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
		return
	}
	//parse json get service name ,addr and port
	service,err := jsonParse.GetRes(string(res))
	if service.Servers == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "decode json error",
		})
		return
	}
	c.String(http.StatusOK, "ok")

	//生成配置文件
	gencnf.GenConf(service)
	//curl dyups api
	curldyups.PostUps(service)
}