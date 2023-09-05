package wire

import "message_queue_service/controller"

var (
	UserController    controller.UserController
	ProductController controller.ProductController
)

func BuildHandler() {
	buildRepo()
	buildService()

	UserController = controller.NewUserController(UserService)
	ProductController = controller.NewProductController(ProductService)
}
