package db

import (
	"gorm.io/gorm"
	"message_queue_service/config"
)

type UserRepo interface {
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepo {
	return &userRepo{
		db: config.MariaDB,
	}
}
