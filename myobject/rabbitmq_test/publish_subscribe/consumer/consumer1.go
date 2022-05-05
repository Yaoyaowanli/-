package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func init() {
	err := os.Setenv("RABBITMQ_IP","amqp://root:yaoyaowanli@172.16.89.206:5672/")
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

// publish_subscribe_test      consumer 1
func main() {
	rabbitMQ()
	defer MQ.Close()

	//创建信道
	ch,err := MQ.Channel()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建队列
	que,err := ch.QueueDeclare("consumer1",
		false,
		false,
		true,
		false,
		nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//声明交换机
	err =ch.ExchangeDeclare("logs","fanout",true,false,false,false,nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//给队列绑定交换机
	err=ch.QueueBind(que.Name,"","logs",false,nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//消费消息
	message,err := ch.Consume(que.Name,"",true,false,false,false,nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//创建channel 堵塞主进程，监听队列
	forever := make(chan bool)
	go func() {
		for d := range message{
			fmt.Println("接收到一条消息:",string(d.Body))
		}
	}()
	fmt.Println("waiting for message ...")
	<-forever
}
