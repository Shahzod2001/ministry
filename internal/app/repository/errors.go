package repository

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrUniverNotFound    = errors.New("university not found")
	ErrAdminNotFound     = errors.New("admin not found")
	ErrPasswordMismatch  = errors.New("password mismatch")
)
