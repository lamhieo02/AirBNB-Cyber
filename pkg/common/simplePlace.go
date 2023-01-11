package common

import (
	"time"
)

// Struct thể hiện thông tin có thể public ra bên ngoài của place
type SimplePlace struct {
	Id            int `json:"-" gorm:"column:id"`
	FakeId        string
	OwnerId       int       `json:"-" gorm:"column:owner_id"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	Name          string    `json:"name" gorm:"column:name"`
	PricePerNight string    `json:"pricePerNight" gorm:"column:price_per_night"`
}

func (SimplePlace) TableName() string {
	return "places"
}
