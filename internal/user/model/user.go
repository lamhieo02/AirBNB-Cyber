package usermodel

import (
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"
)

//type UserRole int

// const (
//
//	RoleGuest UserRole = iota + 1
//	RoleHost
//	RoleAdmin
//
// )
const EntityName = "user"

type User struct {
	common.SqlModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"-" gorm:"column:password"`
	FirstName string        `json:"firstName" gorm:"column:first_name"`
	LastName  string        `json:"lastName" gorm:"column:last_name"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Role      string        `json:"role" gorm:"column:role"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return "users"
}
func (u *User) GetUserId() int {
	return u.Id
}
func (u *User) GetUserEmail() string {
	return u.Email
}

func (u *User) GetUserRole() string {
	return u.Role
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
type UserRegister struct {
	common.SqlModel
	Email     string        `json:"email" gorm:"column:email"`
	Password  string        `json:"password" gorm:"column:password"`
	FirstName string        `json:"firstName" gorm:"column:first_name"`
	LastName  string        `json:"lastName" gorm:"column:last_name"`
	Role      string        `json:"-" gorm:"column:role"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
}

type UserUpdate struct {
	FirstName string        `json:"firstName" gorm:"column:first_name"`
	LastName  string        `json:"lastName" gorm:"column:last_name"`
	Role      string        `json:"role" gorm:"column:role"`
	Phone     string        `json:"phone" gorm:"column:phone"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (u *UserRegister) PrepareCreate() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	// set default role
	u.Role = "guest"
	return nil
}

func (u *UserRegister) Validate() error {
	// business login

	return nil
}
