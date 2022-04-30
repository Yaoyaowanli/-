package  main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

var MQ *amqp.Connection

func init() {
	err := os.Setenv("RABBITMQ_IP","amqp://root:yaoyaowanli@172.16.89.202:5672/")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func rabbitMQ(){
	conn,err:=amqp.Dial(os.Getenv("RABBITMQ_IP"))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	MQ=conn
}

func main() {
	//连接MQ
	rabbitMQ()
	defer MQ.Close()
	//声明通道
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
		nil,
		)
	//读消息,创建消费者
	message,err := ch.Consume(que.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//由于消费者端需要一直监听，所以我们要用一个for循环+channel去阻塞主进程，使得主进程一直处于监听状态。
	forever := make(chan bool)
	go func() {
		for d := range message{
			fmt.Println("收到一条消息：",string(d.Body))
		}
	}()
	log.Println("waiting for message ...")
	<-forever
}
