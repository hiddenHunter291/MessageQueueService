package wire

import "message_queue_service/db"

var (
	UserRepo db.UserRepo
)

func buildRepo() {
	UserRepo = db.NewUserRepo()
}
