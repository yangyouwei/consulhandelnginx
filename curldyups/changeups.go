package curldyups

import (
	"fmt"
	"github.com/yangyouwei/consulhandelnginx/configure"
	"github.com/yangyouwei/consulhandelnginx/jsonParse"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostUps(s *jsonParse.Service)  {
	url := configure.Mainconf.DyupsUrl+s.Name
	responestr := ""
	for _,addr_port := range s.Servers{
		a := "server "+addr_port.Addr+":"+addr_port.Port+";"
		responestr = responestr+a
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(responestr))
	if err != nil {
		fmt.Println(err)
		fmt.Println("add server fail.")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println("add server fail.")
		return
	}

	fmt.Println("dyups repone: ",string(body))
}

