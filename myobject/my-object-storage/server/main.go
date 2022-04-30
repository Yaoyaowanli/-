package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my-object-storage/server/objects"
	"os"
)

func init(){
	err := os.Setenv("LISTEN_ADDRESS",":1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.Setenv("STORAGE_ROOT","/Users/yaoyuan/myobject/my-object-storage")
}

/*func main (){
	//注册给定模式的处理函数Handler
	http.HandleFunc("/test1/",objects.Handler)
	//开启监听
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
}*/

func main () {
	//注册给定模式的处理函数Handler
	engine := gin.Default()
	objects.Post(engine)
	objects.Get(engine)


	err := engine.Run(os.Getenv("LISTEN_ADDRESS"))
	if err != nil {
		log.Println(err)
		return
	}
}
