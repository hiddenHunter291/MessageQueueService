package models

type Product struct {
	ID                 int            `json:"id" gorm:"autoIncrement:primaryKey"`
	ProductName        string         `json:"productName" gorm:"column:product_name;type:varchar(100);NOT NULL"`
	ProductDescription string         `json:"productDescription" gorm:"column:product_description;type:varchar(250);NOT NULL"`
	ProductPrice       int            `json:"productPrice" gorm:"column:product_price;type:bigint(20);NOT NULL"`
	ProductLinks       []ProductLinks `json:"productLinks" gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt          int64          `json:"createdAt" gorm:"column:created_at;type:bigint(20);NOT NULL"`
	UpdatedAt          int64          `json:"updatedAt" gorm:"column:updated_at;type:bigint(20);NOT NULL"`
}

type ProductLinks struct {
	ID        int    `json:"id" gorm:"autoIncrement:primaryKey"`
	ProductID int    `json:"productId" gorm:"product_id"`
	Link      string `json:"link" gorm:"link"`
}

type Order struct {
	ID        int `json:"id" gorm:"autoIncrement:primaryKey"`
	UserID    int `json:"userID" gorm:"column:user_id;type:bigint(20);NOT NULL"`
	ProductID int `json:"productID" gorm:"column:product_id;type:bigint(20);NOT NULL"`
}

type CompressedProductPath struct {
	ID        int    `json:"id" gorm:"autoIncrement:primaryKey"`
	ProductID int    `json:"productId" gorm:"product_id"`
	Path      string `json:"path" gorm:"path"`
}
