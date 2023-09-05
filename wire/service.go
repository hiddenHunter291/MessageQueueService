package wire

import "message_queue_service/service"

var (
	UserService service.UserService
)

func buildService() {
	UserService = service.NewUserService(UserRepo)
}
