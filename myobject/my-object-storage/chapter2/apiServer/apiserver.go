package main

import (
	"log"
	"my-object-storage/chapter2/apiServer/heartBeat"
	"my-object-storage/chapter2/apiServer/locate"
	"my-object-storage/chapter2/apiServer/objects"
	"net/http"
	"os"
)

func main() {
	//开启监听心跳信息
	go heartBeat.ListenHeatBeat()
	http.HandleFunc("/locate/",locate.Handler)
	http.HandleFunc("/objects/",objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}