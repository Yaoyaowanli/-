package main

import (
	"log"
	"my-object-storage/chapter2/dataServer/heartBeat"
	"my-object-storage/chapter2/dataServer/locate"
	"my-object-storage/chapter2/dataServer/objects"
	"net/http"
	"os"
)


func main() {
	//启动一个goroutine进行心跳监听
	go heartBeat.StartHeartBeat()
	//启动一个goroutine来进行locate
	go locate.StartLocate()

	//注册给定模式的处理函数Handler
	http.HandleFunc("/objects/", objects.Handler)
	//开启监听
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil)
	if err != nil{
		log.Fatal(err.Error())
	}
}
