package controller

import "message_queue_service/service"

type ProductController interface {
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}
