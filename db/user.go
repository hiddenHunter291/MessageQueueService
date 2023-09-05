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
	log.Print("In Repo - GetUserId")

	var result models.User
	query := u.db.Debug().
		Table("user").
		Where("id = ?", userId).
		Scan(&result)

	if query.Error != nil {
		return models.User{}, query.Error
	}

	log.Print("Out Repo - GetUserId")
	return result, nil
}

func (u *userRepo) Set(user models.User) error {
	panic("implement me")
}
