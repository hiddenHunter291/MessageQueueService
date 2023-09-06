package wire

import "message_queue_service/service"

var (
	UserService    service.UserService
	ProductService service.ProductService
	FileService    service.FileService
)

func buildService() {
	UserService = service.NewUserService(UserRepo)
	FileService = service.NewFileService()
	ProductService = service.NewProductService(ProductRepo, FileService)
}
