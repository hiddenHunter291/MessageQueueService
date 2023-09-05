package db

import (
	"gorm.io/gorm"
	"message_queue_service/config"
)

type ProductRepo interface {
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo() ProductRepo {
	return &productRepo{
		db: config.MariaDB,
	}
}
