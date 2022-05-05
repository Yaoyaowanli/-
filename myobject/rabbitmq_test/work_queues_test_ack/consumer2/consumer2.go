package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error,msg string){
	if err != nil {
		log.Fatalf("%s:	%s",msg,err)
	}
}


// 学习在work_queues模式下，消息ack、消息持久化、和公平分派
func main() {
	//建立mq连接
	mq,err := amqp.Dial("amqp://root:yaoyaowanli@172.16.89.207:5672/")
	failOnError(err,"dial mq failed : ")
	defer mq.Close()

	//建立信道
	ch,err := mq.Channel()
	failOnError(err,"open channel failed:")
	defer ch.Close()

	//声明队列
	q,err := ch.QueueDeclare("test_ack",true,false,false,false,nil)
	failOnError(err,"declare queue failed:")

	//保证消息公平分发，不至于某一个consumer一直处理，而其他不处理
	err = ch.Qos(
		1,	// prefetch count  在server收到consumer的ACK之前，预取的数量。为1，表示在没收到consumer的ACK之前，只会为其分发一个消息
		0,	// prefetch size 大于0时，表示在收到consumer确认消息之前，将size个字节保留在网络中
		false,	// global  true:Qos对同一个connection的所有channel有效； false:Qos对同一个channel上的所有consumer有效
	)
	failOnError(err,"Qos failed:")

	//声明消费者                                这里不进行自动应答
	msg,err := ch.Consume(q.Name,"",false,false,false,false,nil)
	//forever阻塞进程
	forever := make(chan bool)
	//启动协程开始读取消息
	go func() {
		for d:= range msg {
			log.Printf("consumer2接收到一条消息：%s",string(d.Body))
			time.Sleep(20*time.Second)  //假设消费者2 业务处理需要20秒
			//当consumer 2 接收到消息，处理的过程中G了，那么MQserver接收不到应答，就会把这条消息重新放入队列，分配给或者的consumer
			d.Ack(false) //手动ack
		}
	}()
	log.Println("consumer2(20s) waiting for messages ...")
	<-forever
}