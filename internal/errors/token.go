package errors

import "errors"

var (
	InvalidToken          = errors.New("invalid token")
	FailedTokenGeneration = errors.New("failed token generation")
	UnexpectedTokenMethod = errors.New("unexpected signing method")
)
