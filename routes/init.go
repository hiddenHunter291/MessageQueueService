package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(server *gin.RouterGroup) {

}

func GetRegisteredRoutes(server *gin.Engine) {
	for _, r := range server.Routes() {
		fmt.Printf("Method: %s | Path: %s\n", r.Method, r.Path)
	}
}
