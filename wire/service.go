package wire

import "message_queue_service/service"

var (
	UserService    service.UserService
	ProductService service.ProductService
)

func buildService() {
	UserService = service.NewUserService(UserRepo)
	ProductService = service.NewProductService(ProductRepo)
}
