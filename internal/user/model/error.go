package usermodel

import (
	"errors"
	"go01-airbnb/pkg/common"
)

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(errors.New("email or password invalid"), "email or password invalid")
	ErrEmailExisted           = common.NewCustomError(errors.New("email existed"), "email existed")
)
