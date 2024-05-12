package entity

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

var (
	ErrResourceNotFound = errors.New("resource not found")
)

var (
	ErrReservationNotAvailable = errors.New("reservation not available")
	ErrInvalidReservationTime  = errors.New("invalid reservation time")
)
