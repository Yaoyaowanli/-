package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
)


func init (){
	os.Setenv("RABBITMQ_IP","amqp://root:yaoyaowanli@172.16.89.204:5672/")
}

//rabbitMQ 连接到中间件
func rabbitMQ(){
	conn,err := amqp.Dial(os.Getenv("RABBITMQ_IP"))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	MQ=conn
}

//MQ rabbitMQ 连接
var MQ *amqp.Connection

// rabbitMQ work queues 模式 生产消息
func main() {
	//创建MQ连接
	rabbitMQ()
	defer MQ.Close()
	//创建信道
	ch,err := MQ.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//创建队列
	que,err := ch.QueueDeclare("hello",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//准备消息
	message := "hello world ~"
	for i:=1;i<11;i++{
		msg := strconv.Itoa(i)
		err :=ch.Publish("",
			que.Name,
			false,
			false,
			amqp.Publishing{ContentType: "text/plain",
				Body: []byte(message+msg),   //消息体
			})
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}

}
