package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
)

func init() {
	err:= os.Setenv("RABBITMQ_IP","amqp://root:yaoyaowanli@172.16.89.206:5672/")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}


func rabbitMQ (){
	conn,err := amqp.Dial(os.Getenv("RABBITMQ_IP"))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	MQ=conn
}


var MQ *amqp.Connection

// rabbitMQ  publish/subscribe 模型  （发布/订阅模型）
func main() {
	rabbitMQ()
	defer MQ.Close()

	//声明信道
	ch,err := MQ.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//声明队列
	/*que,err :=ch.QueueDeclare("hello",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}*/
	//声明交换机,在fanout 广播模式下，生产者只能发送消息给交换机，由交换机来决定发送给那个队列，生产者无法决定
	//交换机把消息广播给所有绑定过的队列，队列的消费者都能拿到消息，实现了一个消息被多个消费者所消费。
	err = ch.ExchangeDeclare("logs",	//name是交换机的名称
		"fanout",						//kind 为交换机的类型，fanout就是广播
		true,						//是否持久化
		false,					//是否自动删除
		false,						//是否为内部消息
		false,						//是否阻塞
		nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//准备消息
	msg := "hello world~ "
	for i:=0;i<10;i++ {
		no := strconv.Itoa(i+1)
		err = ch.Publish("logs",
			"fanout",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body: 	[]byte(msg+no),
			})
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	fmt.Println("消息发送完毕~")
}
