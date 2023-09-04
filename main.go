package main

import (
	"github.com/gin-gonic/gin"
	"message_queue_service/routes"
	"net/http"
)

func main() {
	server := gin.New()
	group := server.Group("/api/ecommerce")

	routes.InitializeRoutes(group)

	// health check URL
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "pong"})
	})

	if err := server.Run(":5002"); err != nil {
		panic(err)
	}
}
