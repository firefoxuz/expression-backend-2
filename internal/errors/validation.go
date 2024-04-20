package errors

import "errors"

var (
	LoginMustBeAtLeastFive    = errors.New("login must be at least 5 characters")
	PasswordMustBeAtLeastFive = errors.New("password must be at least 5 characters")
	InvalidCredentials        = errors.New("login or password is incorrect")
)
