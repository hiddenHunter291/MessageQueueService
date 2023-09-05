package db

import (
	"gorm.io/gorm"
	"log"
	"message_queue_service/config"
	"message_queue_service/models"
)

type UserRepo interface {
	GetUserId(userId int) (models.User, error)
	Set(user models.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepo {
	return &userRepo{
		db: config.MariaDB,
	}
}

func (u *userRepo) GetUserId(userId int) (models.User, error) {
	log.Print("In Repo - function GetUserId")

	var result models.User
	query := u.db.Debug().
		Table("users").
		Where("id = ?", userId).
		Scan(&result)

	if query.Error != nil {
		return models.User{}, query.Error
	}

	log.Print("Out Repo - function GetUserId")
	return result, nil
}

func (u *userRepo) Set(user models.User) error {
	log.Print("In Repo - function Set")

	query := u.db.Debug().
		Table("users").
		Create(&user)

	if query.Error != nil {
		return query.Error
	}

	log.Print("Out Repo - function Set")
	return nil
}
