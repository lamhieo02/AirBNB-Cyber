package reviewmodel

import "go01-airbnb/pkg/common"

const EntityName = "Booking"

type Review struct {
	common.SqlModel
	BookingId int    `json:"bookingId" gorm:"column:booking_id"`
	Rating    int    `json:"rating" gorm:"column:rating"`
	Comment   string `json:"comment" gorm:"column:comment"`
}

func (Review) TableName() string {
	return "reviews"
}
