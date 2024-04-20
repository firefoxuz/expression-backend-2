package errors

import "errors"

var (
	CannotConnectDatabase = errors.New("cannot connect database")
	CannotFindEntity      = errors.New("cannot find entity")
	UserAlreadyExists     = errors.New("user already exists")
	UserNotExists         = errors.New("user not exists")
)
