package service

import "message_queue_service/db"

type ProductService interface {
}

type productService struct {
	productRepo db.ProductRepo
}

func NewProductService(productRepo db.ProductRepo) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}
