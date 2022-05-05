package heartBeat

import (
	"my-object-storage/src/rabbitmq"
	"os"
	"time"
)


//StartHeartBeat 创建rabbitMQ结构体每隔5s 向apiServers交换机发送心跳信息，消息体为监听地址
func StartHeartBeat(){
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	for {
		mq.Publish("apiServers",os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5*time.Second)
	}
}
