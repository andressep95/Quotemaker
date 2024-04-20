package errs

import "errors"

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrQuoteNotFound    = errors.New("quote not found")
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidInput     = errors.New("invalid input")
	ErrDatabaseError    = errors.New("database error")
)
