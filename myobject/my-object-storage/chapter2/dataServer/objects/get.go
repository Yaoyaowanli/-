package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)


//get 方法处理http的get请求，在ROOT+objects目录下找到目标文件写入响应w
func get (w http.ResponseWriter,r *http.Request){
	//get环境变量root（项目路径）+ objects（对象存储的文件夹）+ url路径以/分割为字符串切片他的第三个元素就是要get的对象文件的名字
	file,err := os.Open(os.Getenv("STORAGE_ROOT")+"objects"+
		strings.Split(r.URL.EscapedPath(),"/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(w,file)
}
