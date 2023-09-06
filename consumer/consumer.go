package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"message_queue_service/config"
)

func StartConsumer() {
	log.Print("Consumer started listening...")

	run := true
	for run {
		event := config.KafkaConsumer.Poll(100)
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *kafka.Message:
			fmt.Printf("Received kafka event | Event : %s\n", e.String())
			fmt.Printf("Key recieved : %s\n", string(e.Key))
		case kafka.Error:
			fmt.Printf("kafka Error | Code : %v | Error : %v\n", e.Code(), e.Error())
			if e.Code() == kafka.ErrAllBrokersDown {
				run = false
			}
		default:
			fmt.Printf("Ingnored event | Event : %v\n", e)
		}
	}

	log.Print("Consumer stopped listening...")
}
