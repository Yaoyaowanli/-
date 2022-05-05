package objects

import "net/http"


//Handler 分析http请求，进入相应的业务函数
func Handler (w http.ResponseWriter, r *http.Request){
	//分析http请求进入对应的处理业务
	method := r.Method
	if method==http.MethodGet{
		//处理get请求
		get(w,r)
		return
	}else if method==http.MethodPut{
		//处理put请求
		put(w,r)
		return
	}
	//如果method不是get和put，写入404到响应
	w.WriteHeader(http.StatusNotFound)
}
