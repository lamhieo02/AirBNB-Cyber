package locationmodel

import "go01-airbnb/pkg/common"

const EntityName = "Location"

type Location struct {
	common.SqlModel
	Country string `json:"Country" gorm:"column:country"`
	State   string `json:"state" gorm:"column:state"`
	City    string `json:"city" gorm:"column:city"`
}

func (Location) TableName() string {
	return "locations"
}
