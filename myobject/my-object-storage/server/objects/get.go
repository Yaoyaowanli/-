package objects

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Get(engine *gin.Engine){
	engine.GET("/test/", func(context *gin.Context) {
		//处理get请求，取出存储的信息返回
		name := context.PostForm("name")
		file,err := os.Open(os.Getenv("STORAGE_ROOT")+"/test1/"+name+".txt")
		if err != nil {
			log.Println(err)
			context.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		defer file.Close()
		io.Copy(context.Writer,file)
	})
}

/*   使用环境变量
func get (w http.ResponseWriter,r *http.Request){
	//处理get请求，打开文件，把内容写入到w
	file,err := os.Open(os.Getenv("STORAGE_ROOT"+"/objects/"+strings.Split(r.URL.EscapedPath(),"/")[2]))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer file.Close()
	io.Copy(w,file)
}
*/
