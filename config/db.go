package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var MariaDB *gorm.DB

func InitializeDB() {
	db, err := gorm.Open(mysql.Open(Env.MariaDBConn+fmt.Sprintf("?%s", "&parseTime=True")), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true, TablePrefix: "ecommerce."},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 15)

	MariaDB = db
}
