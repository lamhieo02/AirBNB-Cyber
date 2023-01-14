package common

const (
	DBTypePlace          = 1
	DBTypeUser           = 2
	DBTypeAmenity        = 3
	DBPaging             = 4
	DBTypePlaceAmenities = 5
	DBBooking            = 6
	DBReview             = 7
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
