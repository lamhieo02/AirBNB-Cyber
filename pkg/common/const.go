package common

const (
	DBTypePlace = 1
	DBTypeUser  = 2
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
}
