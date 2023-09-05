package routes

import (
	"github.com/gin-gonic/gin"
	"message_queue_service/wire"
)

func createUserRoutes(server *gin.RouterGroup) {
	user := server.Group("/user")

	// user routes
	user.GET("/:id", wire.UserController.Get)
	user.POST("", wire.UserController.Set)
}
