package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var RecordNotFound = errors.New("record not found")

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"` // Giu err goc
	Key        string `json:"key"` // custom code
}

func newErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}

}

//Login , Token

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return newErrorResponse(root, msg, root.Error(), key)
	}
	return newErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return newErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete %s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate %s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet %s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDeletedt %s", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err %s AlreadyExists", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err %s NotFound", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s ", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate %s", entity),
	)
}

func ErrNoPermission(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You have no permission %s ", strings.ToLower(entity)),
		fmt.Sprintf("ErrPermission %s", entity),
	)
}
