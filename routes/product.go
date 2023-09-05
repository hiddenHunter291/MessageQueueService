package routes

import (
	"github.com/gin-gonic/gin"
	"message_queue_service/wire"
)

func createProductRoutes(server *gin.RouterGroup) {
	product := server.Group("/product")

	product.POST("", wire.ProductController.Create)
}
