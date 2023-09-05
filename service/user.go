package service

import (
	"log"
	"message_queue_service/db"
	"message_queue_service/models"
)

type UserService interface {
	Get(userId int) (models.User, error)
	Set(user models.User) error
}

type userService struct {
	userRepo db.UserRepo
}

func NewUserService(userRepo db.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Get(userId int) (models.User, error) {
	log.Print("In Service - Get")

	result, err := u.userRepo.GetUserId(userId)
	if err != nil {
		return models.User{}, err
	}

	log.Print("Out Service - Get")
	return result, nil
}

func (u *userService) Set(user models.User) error {
	panic("implement me")
}
