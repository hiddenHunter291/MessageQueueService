package models

type User struct {
	ID           int     `json:"id" gorm:"autoIncrement:primaryKey"`
	Name         string  `json:"name" gorm:"column:name;type:varchar(100);NOT NULL"`
	MobileNumber int     `json:"mobileNumber" gorm:"column:mobile_number;type:bigint(20);NOT NULL"`
	Latitude     float64 `json:"latitude" gorm:"column:latitude;type:double(10,4);NOT NULL"`
	Longitude    float64 `json:"longitude" gorm:"column:longitude;type:double(10,4);NOT NULL"`
	CreatedAt    int64   `json:"createdAt" gorm:"column:created_at;type:bigint(20);NOT NULL"`
	UpdatedAt    int64   `json:"updatedAt" gorm:"column:updated_at;type:bigint(20);NOT NULL"`
}

type OrderInfo struct {
	UserID             int      `json:"userID"`
	ProductName        string   `json:"productName"`
	ProductDescription string   `json:"productDescription"`
	ProductLinks       []string `json:"productLinks"`
	ProductPrice       int      `json:"productPrice"`
}
