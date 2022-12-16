package common

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`

	Causes error `json:"-"` // root error
	//Key    string `json:"key"` // using to do multi language
}

func (e *AppError) RootErr() error {
	if err, ok := e.Causes.(*AppError); ok {
		return err.RootErr()
	}
	return e.Causes
}

func (e *AppError) Error() string {
	return e.RootErr().Error()
}

func ErrAuthorized(causes error) *AppError {
	return &AppError{Causes: causes, Message: "you have no authorize", StatusCode: http.StatusUnauthorized}
}

func ErrForbidden(causes error) *AppError {
	return &AppError{Causes: causes, Message: "you have no permission", StatusCode: http.StatusForbidden}
}

func ErrBadRequest(causes error) *AppError {
	return &AppError{Causes: causes, Message: "invalid request", StatusCode: http.StatusBadRequest}
}

func NewNotFoundError(causes error) *AppError {
	return &AppError{Causes: causes, Message: "not found", StatusCode: http.StatusNotFound}
}

func ErrInternal(causes error) *AppError {
	return &AppError{Causes: causes, Message: "Something went wrong in the server", StatusCode: http.StatusInternalServerError}
}

func ErrorDB(causes error) *AppError {
	return &AppError{Causes: causes, Message: "Something went wrong with db", StatusCode: http.StatusInternalServerError}
}

func NewCustomError(causes error, msg string) *AppError {
	return &AppError{Causes: causes, Message: msg, StatusCode: http.StatusBadRequest}
}

func ErrCannotListEntity(entity string, cause error) *AppError {
	return NewCustomError(cause, fmt.Sprintf("Cannot list %s", entity))
}

func ErrEntityNotFound(entity string, cause error) *AppError {
	return NewCustomError(cause, fmt.Sprintf("%s not found", entity))
}

func ErrCannotCreateEntity(entity string, cause error) *AppError {
	return NewCustomError(cause, fmt.Sprintf("Cannot create %s", entity))
}

func ErrCannotUpdateEntity(entity string, cause error) *AppError {
	return NewCustomError(cause, fmt.Sprintf("Cannot update %s", entity))
}

func ErrCannotDeleteEntity(entity string, cause error) *AppError {
	return NewCustomError(cause, fmt.Sprintf("Cannot delete %s", entity))
}

func ErrEntityExisted(entity string, cause error) *AppError {
	return NewCustomError(cause, fmt.Sprintf("%s have existed", entity))
}
