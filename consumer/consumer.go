package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"message_queue_service/config"
)

const (
	totalWorkers = 5
)

func spawnWorkers(jobs chan string, numOfWorkers int) {
	for i := 0; i < numOfWorkers; i++ {
		go worker(jobs, i)
	}
}

func StartConsumer() {
	log.Print("Consumer started listening...")

	jobs := make(chan string, 5)
	spawnWorkers(jobs, totalWorkers)

	run := true
	for run {
		event := config.KafkaConsumer.Poll(100)
		if event == nil {
			continue
		}

		switch e := event.(type) {
		case *kafka.Message:
			fmt.Printf("Received kafka event | Event : %s | Key : %s\n", e.String(), string(e.Key))
			jobs <- string(e.Key)

		case kafka.Error:
			fmt.Printf("kafka Error | Code : %v | Error : %v\n", e.Code(), e.Error())
			if e.Code() == kafka.ErrAllBrokersDown {
				run = false
				close(jobs)
			}

		default:
			fmt.Printf("Ingnored event | Event : %v\n", e)
		}
	}

	log.Print("Consumer stopped listening...")
}
