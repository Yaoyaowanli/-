package locate

import (
	"my-object-storage/src/rabbitmq"
	"os"
	"strconv"
)


//locate 检查name文件是否存在
func locate (name string)bool{
	//使用stat寻找文件
	_,err := os.Stat(name)
	//isNotExist 根据错误判断此文件是否不存在，返回bool
	return ! os.IsNotExist(err)
}

//StartLocate  开启定位服务，监听接口服务发来的get请求,接收到get请求调用locate函数在本地寻找文件对象是否存在本节点，如果存在
//调用rabbitMQ.send方法发送本节点地址给接口服务的队列
func StartLocate (){
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	//绑定dataServers交换机
	mq.Bind("dataServers")
	//消费消息
	ch := mq.Consume()
	for msg:= range ch{
		//去除双引号返回需要get的文件名
		object,err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		//拼接完整路径
		str := locate(os.Getenv("STORAGE_ROOT")+"/objects/"+object)
		if str {
			//文件存在，send本节点的地址告诉接口服务要找的对象在这
			mq.Send(msg.ReplyTo,os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
