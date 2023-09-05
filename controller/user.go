package controller

import (
	"github.com/gin-gonic/gin"
	"message_queue_service/service"
)

type UserController interface {
	Get(c *gin.Context)
	Set(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u userController) Get(c *gin.Context) {
	panic("implement me")
}

func (u userController) Set(c *gin.Context) {
	panic("implement me")
}
