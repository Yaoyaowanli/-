package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func failOnError(err error,msg string){
	if err != nil {
		log.Fatalf("%s:	%s",msg,err)
	}
}


// 学习在work_queues模式下，消息ack、消息持久化、和公平分派
func main() {
	mq,err := amqp.Dial("amqp://root:yaoyaowanli@172.16.89.207:5672/")
	failOnError(err,"dial rabbitmq failed :")
	defer mq.Close()

	//创建信道
	ch,err := mq.Channel()
	failOnError(err,"open channel failed :")
	defer ch.Close()


	//创建队列
	//打开队列的持久化，我们在保证消息持久化之前应该先让队列持久化
	q,err := ch.QueueDeclare("test_ack",true,false,false,false,nil)
	failOnError(err,"declare queue failed :")

	//队列持久化了，下来准备消息，并将消息持久化
	msg := "hello ~"
	for i:=0;i<10;i++{
		n:=strconv.Itoa(i+1)
		err:=ch.Publish("",q.Name,false,false,amqp.Publishing{
			ContentType: "text/plain",
			DeliveryMode: amqp.Persistent, //deliverymode 交付模式 persistent 持久
			Body: []byte(msg+n),
		})
		failOnError(err,"publish failed :")
		log.Printf("消息%d发送成功!\n",i+1)
	}
}