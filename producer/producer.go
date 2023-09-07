package producer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"message_queue_service/config"
	"strconv"
)

type Producer interface {
	ProduceEvent(key int) error
}

type producer struct {
	kafka *kafka.Producer
	topic string
}

func NewKafkaProducer(topic string) Producer {
	return &producer{
		kafka: config.KafkaProducer,
		topic: topic,
	}
}

func (p *producer) ProduceEvent(key int) error {
	log.Print("In producer - function ProduceEvent")

	strKey := strconv.Itoa(key)
	err := p.kafka.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Key:            []byte(strKey),
	}, nil)

	if err != nil {
		return err
	}

	log.Print("Out producer - function ProduceEvent")
	return nil
}
