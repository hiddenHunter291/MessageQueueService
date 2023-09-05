package db

import (
	"gorm.io/gorm"
	"log"
	"message_queue_service/config"
	"message_queue_service/models"
)

type ProductRepo interface {
	CreateProduct(product models.Product) error
	CreateOrder(userID int, productName, productDescription string) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo() ProductRepo {
	return &productRepo{
		db: config.MariaDB,
	}
}

func (p *productRepo) CreateProduct(product models.Product) error {
	log.Print("In Repo - function CreateProduct")

	query := p.db.Debug().
		Table("products").
		Create(&product)

	if query.Error != nil {
		return query.Error
	}

	log.Print("Out Repo - function CreateProduct")
	return nil
}

func (p *productRepo) CreateOrder(userID int, productName, productDescription string) error {
	log.Print("In Repo - function CreateOrder")

	var (
		productID int
		order     models.Order
	)

	query := p.db.Debug().
		Table("products").
		Select("id").
		Where("product_name = ? and product_description = ?", productName, productDescription).
		Scan(&productID)

	if query.Error != nil {
		return query.Error
	}

	order.ProductID = productID
	order.UserID = userID

	query = p.db.Debug().
		Table("orders").
		Create(&order)

	if query.Error != nil {
		return query.Error
	}

	log.Print("Out Repo - function CreateOrder")
	return nil
}
