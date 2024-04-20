package service

import (
	constErrors "expression-backend/internal/errors"
	"strings"
	"unicode/utf8"
)

const minimumCredentialLength = 5

func trimCredential(credential string) string {
	return strings.TrimSpace(credential)
}

func ValidateLogin(credential string) error {
	if utf8.RuneCountInString(trimCredential(credential)) < minimumCredentialLength {
		return constErrors.LoginMustBeAtLeastFive
	}
	return nil
}

func ValidatePassword(credential string) error {
	if utf8.RuneCountInString(trimCredential(credential)) < minimumCredentialLength {
		return constErrors.PasswordMustBeAtLeastFive
	}
	return nil
}
