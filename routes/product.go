package routes

import "github.com/gin-gonic/gin"

func createProductRoutes(server *gin.RouterGroup) {
	product := server.Group("/product")

	product.POST("")
}
