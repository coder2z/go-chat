package handler

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/streadway/amqp"
	"socket/socket-web/config"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	key string
	//连接信息
	MqUrl string
}

//创建RabbitMQ结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitMq := &RabbitMQ{QueueName: queueName, Exchange: exchange, key: key, MqUrl: fmt.Sprintf("amqp://%s:%s@%s",
		config.RabbitMqConfig.User,
		config.RabbitMqConfig.Password,
		config.RabbitMqConfig.Host)}
	var err error
	//创建RabbitMQ连接
	rabbitMq.conn, err = amqp.Dial(rabbitMq.MqUrl)
	rabbitMq.failOnErr(err, "创建连接错误!")
	rabbitMq.channel, err = rabbitMq.conn.Channel()
	rabbitMq.failOnErr(err, "获取channel失败!")
	return rabbitMq
}

//断开channel和connection
func (r *RabbitMQ) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
	}
}

//订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//订阅模式消费端代码
func (r *RabbitMQ) RecieveSub() {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil)
	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range messges {
			for key, value := range ConnAll {
				if value {
					_ = key.WriteMessage(d.Body)
				}
			}
		}
	}()
	fmt.Println("mq消费启动")
	<-forever
}
