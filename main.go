package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message_queue_service/config"
	"message_queue_service/routes"
	"net/http"
)

func main() {
	fmt.Println("--------------------Starting Ecommerce Service--------------------")
	config.InitializeEnv()
	config.InitializeDB()

	server := gin.New()
	group := server.Group("/api/ecommerce")

	routes.InitializeRoutes(group)
	routes.GetRegisteredRoutes(server)

	// health check URL
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "alive"})
	})

	if err := server.Run(":5002"); err != nil {
		panic(err)
	}
}
