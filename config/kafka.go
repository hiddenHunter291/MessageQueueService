package config

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	KafkaConsumer *kafka.Consumer
	KafkaProducer *kafka.Producer
)

func InitializeConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": Env.KafkaServer,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{Env.KafkaTopic}, nil)
	if err != nil {
		panic(err)
	}

	KafkaConsumer = c
	fmt.Println("Kafka consumer initialized")
}

func InitializeProducer() {
	k, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": Env.KafkaServer,
	})

	if err != nil {
		panic(err)
	}

	KafkaProducer = k
	fmt.Println("Kafka producer initialized")
}
