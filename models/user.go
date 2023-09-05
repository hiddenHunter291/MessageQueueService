package models

type User struct {
	Id           int     `json:"id" gorm:"id"`
	Name         string  `json:"name" gorm:"name"`
	MobileNumber int     `json:"mobileNumber" gorm:"mobile_number"`
	Latitude     float64 `json:"latitude" gorm:"latitude"`
	Longitude    float64 `json:"longitude" gorm:"longitude"`
	CreatedAt    int64   `json:"createdAt" gorm:"create_at"`
	UpdatedAt    int64   `json:"updatedAt" gorm:"updated_at"`
}
