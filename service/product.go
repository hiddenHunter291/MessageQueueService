package service

import (
	"log"
	"message_queue_service/db"
	"message_queue_service/models"
	"time"
)

type ProductService interface {
	Create(order models.OrderInfo) error
}

type productService struct {
	productRepo db.ProductRepo
}

func NewProductService(productRepo db.ProductRepo) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p *productService) Create(order models.OrderInfo) error {
	log.Print("In service - function Create")

	var err error
	productLinks := make([]models.ProductLinks, 0)

	for _, link := range order.ProductLinks {
		productLinks = append(productLinks, models.ProductLinks{Link: link})
	}

	product := models.Product{
		ProductName:        order.ProductName,
		ProductDescription: order.ProductDescription,
		ProductPrice:       order.ProductPrice,
		ProductLinks:       productLinks,
		CreatedAt:          time.Now().UnixMilli(),
		UpdatedAt:          time.Now().UnixMilli(),
	}

	err = p.productRepo.CreateProduct(product)
	if err != nil {
		return err
	}

	err = p.productRepo.CreateOrder(order.UserID, product.ProductName, product.ProductDescription)
	if err != nil {
		return err
	}

	log.Print("Out service - function Create")
	return nil
}
