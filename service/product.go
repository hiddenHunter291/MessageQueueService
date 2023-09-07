package service

import (
	"log"
	"message_queue_service/db"
	"message_queue_service/models"
	"message_queue_service/producer"
	"time"
)

type ProductService interface {
	Create(order models.OrderInfo) error
}

type productService struct {
	productRepo   db.ProductRepo
	kafkaProducer producer.Producer
}

func NewProductService(productRepo db.ProductRepo, kafkaProducer producer.Producer) ProductService {
	return &productService{
		productRepo:   productRepo,
		kafkaProducer: kafkaProducer,
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
		CreatedAt:          time.Now().Unix(),
		UpdatedAt:          time.Now().Unix(),
	}

	err = p.productRepo.CreateProduct(product)
	if err != nil {
		return err
	}

	err = p.productRepo.CreateOrder(order.UserID, product.ProductName, product.ProductDescription)
	if err != nil {
		return err
	}

	productID, getErr := p.productRepo.GetProductID(order.ProductName, order.ProductDescription)
	if getErr != nil {
		return getErr
	}

	kfkErr := p.kafkaProducer.ProduceEvent(productID)

	if kfkErr != nil {
		return kfkErr
	}

	log.Print("Out service - function Create")
	return nil
}
