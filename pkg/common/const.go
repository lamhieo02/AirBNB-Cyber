package common

const (
	DBTypePlace   = 1
	DBTypeUser    = 2
	DBTypeAmenity = 3
	DBPaging      = 4
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
