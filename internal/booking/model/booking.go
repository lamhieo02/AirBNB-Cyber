package bookingmodel

import (
	"go01-airbnb/pkg/common"
)

type Booking struct {
	common.SqlModel
	UserId       int                 `json:"-" gorm:"column:user_id"`
	User         *common.SimpleUser  `json:"user" gorm:"preload:false"`
	PlaceId      int                 `json:"-" gorm:"column:place_id"`
	Place        *common.SimplePlace `json:"place" gorm:"preload:false"`
	Status       string              `json:"status" gorm:"column:status"`
	CheckInDate  string              `json:"check_in_date" gorm:"column:checkin_date"`
	CheckOutDate string              `json:"check_out_date" gorm:"column:checkout_date"`
}

func (Booking) TableName() string {
	return "bookings"
}

const EntityName = "Bookings"
