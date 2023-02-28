package common

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
)

type AppError struct {
	StatusCode codes.Code `json:"status_code"`
	RootErr    error      `json:"-"`
	Message    string     `json:"message"`
	Log        string     `json:"log"`
	Key        string     `json:"error_key"`
}

func NewErrorBadRequest(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: codes.InvalidArgument,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode codes.Code, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: codes.Unauthenticated,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorBadRequest(root, msg, root.Error(), key)
	}

	return NewErrorBadRequest(errors.New(msg), msg, msg, key)
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
	return NewFullErrorResponse(codes.Internal, err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorBadRequest(err, "Invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(codes.Internal, err,
		"Something went wrong in the server", err.Error(), "ErrInternal")
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
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyExists", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		"You have no permission",
		"ErrNoPermission",
	)
}

func ErrRequestCanceled() *AppError {
	return NewFullErrorResponse(
		codes.Canceled,
		errors.New("request is canceled"),
		"Request canceled!",
		"request is canceled",
		"ErrRequestCanceled",
	)
}

func ErrDeadlineExceeded() *AppError {
	return NewFullErrorResponse(
		codes.DeadlineExceeded,
		errors.New("deadline is exceeded"),
		"Deadline is exceeded!",
		"deadline is exceeded",
		"ErrDeadlineExceeded",
	)
}

var ErrWrongUID = NewCustomError(
	errors.New("wrong uid"),
	"Wrong uid",
	"ErrWrongUID",
)

var (
	RecordNotFound = NewFullErrorResponse(
		http.StatusNotFound,
		errors.New("record not found"),
		"Record not found",
		errors.New("record not found").Error(),
		"RecordNotFound",
	)
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
