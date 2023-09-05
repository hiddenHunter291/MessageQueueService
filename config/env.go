package config

import (
	"fmt"
	"os"
)

type envData struct {
	MariaDBConn string
}

var Env *envData

func InitializeEnv() {
	Env = &envData{
		MariaDBConn: os.Getenv("databaseConnection"),
	}
	fmt.Println("Envs Initialized successfully")
}
