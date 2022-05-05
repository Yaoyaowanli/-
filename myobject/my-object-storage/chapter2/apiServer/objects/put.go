package objects

import (
	"log"
	"net/http"
	"strings"
)

//put 将http的put请求转发给数据服务
func put (w http.ResponseWriter ,r *http.Request){
	object := strings.Split(r.URL.EscapedPath(),"/")[2]
	c,err := storeObject(r.Body,object)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(c)
}