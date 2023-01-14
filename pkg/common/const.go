package common

const (
	DBTypePlace          = 1
	DBTypeUser           = 2
	DBTypeAmenity        = 3
	DBPaging             = 4
	DBTypePlaceAmenities = 5
	DBTypeBooking        = 6
	DBTypeReview         = 7
	DBTypeLocation       = 8
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
