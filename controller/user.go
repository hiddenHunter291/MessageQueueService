package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"message_queue_service/models"
	"message_queue_service/service"
	"net/http"
	"strconv"
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

func (u *userController) Get(c *gin.Context) {
	log.Print("In controller - function Get")

	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, servErr := u.userService.Get(int(id))
	if servErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": servErr.Error()})
		return
	}

	log.Print("Out controller - function Get")
	c.JSON(http.StatusOK, gin.H{"user": result})
}

func (u *userController) Set(c *gin.Context) {
	log.Print("In controller - function Set")

	var input models.User
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	servErr := u.userService.Set(input)
	if servErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": servErr.Error()})
		return
	}

	log.Print("Out controller - function Set")
	c.JSON(http.StatusOK, gin.H{"msg": "user created successfully"})
}
