package service

import (
	"github.com/streadway/amqp"
	"log"
)

var (
	exchangeName = "hello-exchange"
	queueName    = "hello-queue"
	routingKey   = "hello-routing-key"
)

type RabbitService struct {
	Rabbit *amqp.Connection
}

func (iService RabbitService) InitQueue() {
	// 创建信道
	ch, err := iService.Rabbit.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()
	// 声明exchange
	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeDirect,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}
	// 声明queue
	_, err = ch.QueueDeclare(
		queueName,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}
	// 绑定队列
	err = ch.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %s", err)
	}
}

func (iService RabbitService) Send(msg string) error {
	ch, err := iService.Rabbit.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()
	// 消息
	err = ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	return err
}

func (iService RabbitService) Recv() {
	ch, err := iService.Rabbit.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
