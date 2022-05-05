package locate

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter,r *http.Request){
	m := r.Method
	//如果请求不是get
	if m != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//查找数据服务节点取到目标文件所在的节点地址
	info := Locate(strings.Split(r.URL.EscapedPath(),"/")[2])
	if len(info)==0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b,_ := json.Marshal(info)
	w.Write(b)
}
