package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

type envData struct {
	MariaDBConn string `validate:"required"`
	KafkaServer string `validate:"required"`
	KafkaTopic  string `validate:"required"`
	KafkaGroup  string `validate:"required"`
}

var Env *envData

func InitializeEnv() {
	Env = &envData{
		MariaDBConn: os.Getenv("databaseConnection"),
		KafkaServer: os.Getenv("kafkaServer"),
		KafkaTopic:  os.Getenv("kafkaTopic"),
		KafkaGroup:  os.Getenv("kafkaGroup"),
	}

	validate := validator.New()
	if err := validate.Struct(Env); err != nil {
		panic(err)
	}

	fmt.Println("Envs Initialized successfully")
}
