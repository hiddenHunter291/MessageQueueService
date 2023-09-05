package wire

import "message_queue_service/db"

var (
	UserRepo    db.UserRepo
	ProductRepo db.ProductRepo
)

func buildRepo() {
	UserRepo = db.NewUserRepo()
	ProductRepo = db.NewProductRepo()
}
