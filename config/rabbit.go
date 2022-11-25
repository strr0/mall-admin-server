package config

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var (
	conn *amqp.Connection
)

type RabbitConfig struct {
	Url      string
	Username string
	Password string
}

func InitRabbit() {
	var err error
	var rabbitConfig RabbitConfig
	err = settings.UnmarshalKey("rabbit", &rabbitConfig)
	if err != nil {
		log.Fatalf("get config failed: %v", err)
	}
	// 建立连接
	url := fmt.Sprintf("amqp://%s:%s@%s", rabbitConfig.Username, rabbitConfig.Password, rabbitConfig.Url)
	conn, err = amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
}

func CloseRabbit() {
	_ = conn.Close()
}

func GetRabbit() *amqp.Connection {
	return conn
}
