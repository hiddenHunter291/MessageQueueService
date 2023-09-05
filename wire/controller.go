package wire

import "message_queue_service/controller"

var (
	UserController controller.UserController
)

func BuildHandler() {
	buildService()
	buildRepo()

	UserController = controller.NewUserController(UserService)
}
