package locate

import (
	"my-object-storage/src/rabbitmq"
	"os"
	"strconv"
	"time"
)



func Locate (name string) string {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	//像所有数据服务节点群发定位name文件对象
	mq.Publish("dataServers",name)
	//接收反馈
	c := mq.Consume()
	//等待1s，如果没有反馈消息说明没有这个文件对象，关闭队列
	go func() {
		time.Sleep(time.Second)
		mq.Close()
	}()
	msg := <- c
	//这里找到了会返回目标文件所在的数据服务节点的地址
	s,_ := strconv.Unquote(string(msg.Body))
	return s
}

//Exist 通过检查Locate结果是否为空字符串来判定对象是否存在
func Exist (name string)bool{
	return Locate(name) != ""
}
