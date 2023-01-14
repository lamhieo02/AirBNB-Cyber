package common

type SimpleReview struct {
	//Id      int       `json:"-" gorm:"column:id"`
	//FakeId  string    `json:"id" gorm:"-"`
	Rating    int    `json:"rating" gorm:"rating"`
	Comment   string `json:"comment" gorm:"comment"`
	PlaceId   int    `json:"-" gorm:"place_id"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Avatar    *Image `json:"avatar" gorm:"column:avatar"`
}
