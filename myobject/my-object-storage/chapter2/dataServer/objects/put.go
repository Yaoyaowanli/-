package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	//                                                   如：127.0.0.1:1234/object/test
	//os.create根据路径创建文件，路径为：环境变量路径+/objects/+对url以/切割拿到他最后的【2】路径
	file, err := os.Create(os.Getenv("STORAGE_ROOT") + "/test1/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		//创建文件失败给客户端返回错误状态码"500：服务端处理错误"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	//把body中的数据拷贝到文件中
	io.Copy(file, r.Body)
}