package common

type SimpleLocation struct {
	Id      int    `json:"-" gorm:"column:id"`
	FakeId  string `json:"id" gorm:"-"`
	Country string `json:"country" gorm:"column:country"`
	State   string `json:"state" gorm:"column:state"`
	City    string `json:"city" gorm:"column:city"`
}

func (SimpleLocation) TableName() string {
	return "locations"
}
