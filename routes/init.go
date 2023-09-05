package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message_queue_service/wire"
)

func InitializeRoutes(server *gin.RouterGroup) {
	wire.BuildHandler()

	createUserRoutes(server)
	createProductRoutes(server)
}

func GetRegisteredRoutes(server *gin.Engine) {
	for _, r := range server.Routes() {
		fmt.Printf("Method: %s | Path: %s\n", r.Method, r.Path)
	}
}
