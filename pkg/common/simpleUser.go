package common

// Struct thể hiện thông tin có thể public ra bên ngoài của user
type SimpleUser struct {
	Id        int    `json:"-" gorm:"column:id"`
	FakeId    string `json:"id" gorm:"-"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
}

func (SimpleUser) TableName() string {
	return "users"
}
