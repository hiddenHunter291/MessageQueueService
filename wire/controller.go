package wire

import "message_queue_service/controller"

var (
	UserController controller.UserController
)

func BuildHandler() {
	buildRepo()
	buildService()

	UserController = controller.NewUserController(UserService)
}
