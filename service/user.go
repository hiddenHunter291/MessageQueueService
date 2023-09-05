package service

import (
	"log"
	"message_queue_service/db"
	"message_queue_service/models"
	"time"
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
	log.Print("In service - function Get")

	result, err := u.userRepo.GetUserId(userId)
	if err != nil {
		return models.User{}, err
	}

	log.Print("Out service - function Get")
	return result, nil
}

func (u *userService) Set(user models.User) error {
	log.Print("In service - function Set")

	user.CreatedAt = time.Now().UnixMilli()
	user.UpdatedAt = time.Now().UnixMilli()

	err := u.userRepo.Set(user)
	if err != nil {
		return err
	}

	log.Print("Out service - function Set")
	return nil
}
