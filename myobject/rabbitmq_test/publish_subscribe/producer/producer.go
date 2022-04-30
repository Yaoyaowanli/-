package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func init() {
	err:= os.Setenv("RABBITMQ_IP","amqp://root:yaoyaowanli@172.16.89.202:5672/")
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

}
