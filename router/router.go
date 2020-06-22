package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yangyouwei/consulhandelnginx/api"
	"github.com/yangyouwei/consulhandelnginx/conflib"
	"github.com/yangyouwei/consulhandelnginx/loglib"
	"github.com/yangyouwei/consulhandelnginx/middleware"
	"net/http"
)

var Router *mux.Router
var mainconf conflib.MainConf

func InitRouter()  {
	mainconf = conflib.Mainconf
	Router = mux.NewRouter()

	//add route
	//Router.HandleFunc("/", api.DoSmoeThing)

	// add static
	//Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	apiroute(Router)

	//use middleware
	//logging
	Router.Use(middleware.LoggingMiddleWare)
	//crossdomain
	//Router.Use(middleware.CrossDomainMiddleWare)

	//loglib.Mylog.Println("server listen on ")
	fmt.Println("server listen on "+ conflib.Mainconf.ListenPort)
	Run(conflib.Mainconf.ListenPort,Router)
}



func Run(p string,router *mux.Router)  {
	err := http.ListenAndServe(":"+p, router)
	if err != nil {
		loglib.Mylog.Panic(err)
		return
	}
}

func apiroute(r *mux.Router)  {
	s := r.PathPrefix("/api").Subrouter()
	//sub api
	s.HandleFunc("/dosmoething0", api.DoSmoeThing)
	s.HandleFunc("/dosmoething1", api.DoSmoeThing)
}