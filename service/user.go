package service

import "message_queue_service/db"

type UserService interface {
}

type userService struct {
	userRepo db.UserRepo
}

func NewUserService(userRepo db.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
