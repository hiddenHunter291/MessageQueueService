package config

import (
	"fmt"
	"os"
)

type envData struct {
	MariaDBConn string
	KafkaServer string
	KafkaTopic  string
}

var Env *envData

func InitializeEnv() {
	Env = &envData{
		MariaDBConn: os.Getenv("databaseConnection"),
		KafkaServer: os.Getenv("kafkaServer"),
		KafkaTopic:  os.Getenv("kafkaTopic"),
	}
	fmt.Println("Envs Initialized successfully")
}
