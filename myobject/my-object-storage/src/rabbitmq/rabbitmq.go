package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

//rabbitMQ 结构体里存放了MQ的连接，MQ的channel，交换机，和队列名
type rabbitMQ struct {
	conn *amqp.Connection  	//MQ连接
	channel *amqp.Channel  	//MQ信道
	QueueName string   			//队列名
	exchange string       	//交换机名
}

//New 建立rabbitMQ 连接
func New(s string) *rabbitMQ {

	//MQ 连接
	conn,err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}

	//建立MQ信道
	ch,err := conn.Channel()
	if err != nil {
		panic(err)
	}

	//声明MQ队列
	que,err := ch.QueueDeclare("",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	mq := new(rabbitMQ)
	mq.conn=conn
	mq.channel=ch
	mq.QueueName=que.Name
	return mq
}


//Bind 绑定MQ交换机模式
func (mq *rabbitMQ)Bind (exchange string){
	err := mq.channel.QueueBind(mq.QueueName,
		"",
		exchange,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	mq.exchange=exchange
}

//Send MQ直接发消息给队列
func (mq *rabbitMQ)Send (queue string,mes interface{}){
	body,err := json.Marshal(mes)
	if err != nil {
		panic(err)
	}
	err = mq.channel.Publish("",
		queue,
		false,
		false,
		amqp.Publishing{Body: body,
			ReplyTo: mq.QueueName,
		},
	)
	if err != nil {
		panic(err)
	}
}

//Publish MQ发布消息到交换机
func (mq *rabbitMQ)Publish(exchange string,mes interface{}){
	body,err := json.Marshal(mes)
	if err != nil {
		panic(err)
	}
	//发送消息给交换机
	err = mq.channel.Publish(exchange,
		"",
		false,
		false,
		amqp.Publishing{
		Body: body,
		ReplyTo: mq.QueueName,
		},
	)
	if err != nil {
		panic(err)
	}
}

//Consume 消费消息,返回一个 <- chan amqp.Delivery
func (mq *rabbitMQ)Consume() <-chan amqp.Delivery{
	d,err := mq.channel.Consume(mq.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return d
}

//Close 关闭conn，channel连接
func (mq *rabbitMQ)Close(){
	mq.channel.Close()
	mq.conn.Close()
}