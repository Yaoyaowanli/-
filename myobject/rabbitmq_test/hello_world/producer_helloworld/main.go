package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

var MQ *amqp.Connection

func init (){
	err := os.Setenv("RABBITMQ_IP","amqp://root:yaoyaowanli@172.16.89.202:5672/")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}

//创建rabbitMQ连接
func rabbitMQ(connStr string) {
	conn,err := amqp.Dial(connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	MQ=conn
}

//使用go连接rabbitmq中的hello world模式来发送和接收消息
func main () {
	rabbitMQ(os.Getenv("RABBITMQ_IP"))
	defer MQ.Close()
//创建信道
	ch,err:= MQ.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	que,err:=ch.QueueDeclare("hello",	//队列名字
		false,   	//是否持久化到磁盘   队列中所有的数据都是在内存中的，如果为true的话，这个通道关闭之后，数据就会存在磁盘中持久化，false的话就会丢弃
		false, 	//不需要用到队列的时候，是否将消息删除
		false,	 	//是否独占队列，true的话，就是只能是这个进程独占这个队列，其他都不能对这个队列进行读写
		false,		//是否阻塞
		nil,			//其他参数
		)
	if err != nil {
		log.Fatal(err.Error())
	}
	//准备开始发送消息
	message := "hello world"

	err = ch.Publish("",	//交换机
		que.Name,					//队列名
		false,			//是否强制性
	// 当mandatory标志位设置为true时，如果exchange根据自身类型和消息routeKey无法找到一个符合条件的queue，那么会调用basic.return方法将消息返回给生产者
	// 当mandatory设置为false时，出现上述情形broker会直接将消息扔掉
		false,			//当immediate标志位设置为true时，如果exchange在将消息路由到queue(s)时发现对于的queue上么有消费者，那么这条消息不会放入队列中。当与消息routeKey关联的所有queue（一个或者多个）都没有消费者时，该消息会通过basic.return方法返还给生产者
	// 是否立刻
	/**
	概括来说，mandatory标志告诉服务器至少将该消息route到一个队列中，否则将消息返还给生产者；immediate标志告诉服务器如果该消息关联的queue上有消费者，则马上将消息投递给它，如果所有queue都没有消费者，直接把消息返还给生产者，不用将消息入队列等待消费者了。
	**/
	amqp.Publishing{
		ContentType: "text/plain",	//传输类型
		Body: []byte(message),		//消息体
	},
	)

	fmt.Println("消息发送成功～")

}
