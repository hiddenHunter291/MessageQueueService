package config

import (
	"fmt"
	"os"
)

type envData struct {
	MariaDBConn string
	KafkaServer string
	KafkaTopic  string
	KafkaGroup  string
}

var Env *envData

func InitializeEnv() {
	Env = &envData{
		MariaDBConn: os.Getenv("databaseConnection"),
		KafkaServer: os.Getenv("kafkaServer"),
		KafkaTopic:  os.Getenv("kafkaTopic"),
		KafkaGroup:  os.Getenv("kafkaGroup"),
	}
	fmt.Println("Envs Initialized successfully")
}
