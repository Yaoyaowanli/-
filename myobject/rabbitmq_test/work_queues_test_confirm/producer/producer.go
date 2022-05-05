package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError (err error,msg string){
	if err != nil{
		log.Println(msg,":",err)
	}
}

//测试rabbitMQ work_queues下发布确认confirm
func main() {
	//连接到MQ
	conn,err := amqp.Dial("amqp://root:yaoyaowanli@172.16.89.207:5672/")
	failOnError(err,"dial MQ failed :")
	defer conn.Close()

	//开启信道
	ch,err := conn.Channel()
	failOnError(err,"open channel failed :")
	defer ch.Close()

	//声明发布确认
	err = ch.Confirm()
	failOnError(err,"confirm failed :")
}
