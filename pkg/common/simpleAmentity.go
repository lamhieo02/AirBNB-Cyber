package common

type SimpleAmenity struct {
	Id          int    `json:"-" gorm:"column:id"`
	FakeId      string `json:"id" gorm:"-"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

func (SimpleAmenity) TableName() string {
	return "amenities"
}
