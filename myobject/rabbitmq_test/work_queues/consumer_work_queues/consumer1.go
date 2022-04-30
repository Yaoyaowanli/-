package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func init (){
	err := os.Setenv("RABBITMQ_IP","amqp://root:yaoyaowanli@172.16.89.204:5672/")
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
	MQ = conn
}

//MQ rabbitMQ 连接
var MQ *amqp.Connection


// MQ work_queues test 消费者1
func main() {
	//连接MQ
	rabbitMQ()
	defer MQ.Close()

	//创建信道
	ch,err := MQ.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建队列
	que,err :=ch.QueueDeclare("hello",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//声明消费者
	message,err:=ch.Consume(que.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//监听，阻塞主进程
	forever := make(chan bool)
	go func() {
		for d := range message {
			fmt.Println("consumer1接收到一条消息:",string(d.Body))
		}
	}()
	fmt.Println("waiting fro message ...")
	<-forever
}
