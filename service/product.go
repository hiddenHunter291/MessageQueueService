package service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"message_queue_service/config"
	"message_queue_service/db"
	"message_queue_service/models"
	"strconv"
	"time"
)

type ProductService interface {
	Create(order models.OrderInfo) error
}

type productService struct {
	fileService FileService
	productRepo db.ProductRepo
}

func NewProductService(productRepo db.ProductRepo, fileService FileService) ProductService {
	return &productService{
		fileService: fileService,
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

	productID, getErr := p.productRepo.GetProductID(order.ProductName, order.ProductDescription)
	if getErr != nil {
		return getErr
	}

	kfkErr := config.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &config.Env.KafkaTopic, Partition: kafka.PartitionAny},
		Key:            []byte(strconv.Itoa(productID)),
	}, nil)

	if kfkErr != nil {
		return kfkErr
	}

	log.Print("Out service - function Create")
	return nil
}
