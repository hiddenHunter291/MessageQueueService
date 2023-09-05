package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"message_queue_service/models"
	"message_queue_service/service"
	"net/http"
)

type ProductController interface {
	Create(c *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (p *productController) Create(c *gin.Context) {
	log.Print("In controller - function Create")

	var input models.OrderInfo
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	servErr := p.productService.Create(input)
	if servErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": servErr.Error()})
		return
	}

	log.Print("Out controller - function Create")
	c.JSON(http.StatusOK, gin.H{"msg": "created successfully"})
}
