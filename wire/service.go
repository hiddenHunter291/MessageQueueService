package wire

import (
	"message_queue_service/config"
	"message_queue_service/producer"
	"message_queue_service/service"
)

var (
	UserService    service.UserService
	ProductService service.ProductService
	FileService    service.FileService
	KafkaProducer  producer.Producer
)

func buildService() {
	KafkaProducer = producer.NewKafkaProducer(config.Env.KafkaTopic)
	UserService = service.NewUserService(UserRepo)
	FileService = service.NewFileService()
	ProductService = service.NewProductService(ProductRepo, KafkaProducer)
}
