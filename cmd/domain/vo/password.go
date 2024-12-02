package vo

import (
	"errors"
	"unicode/utf8"
)

type Password struct {
	Value string
}

func NewPassword(value string) (*Password, error) {
	if utf8.RuneCountInString(value) < 6 {
		return nil, errors.New("invalid password: must be greater than 5")
	}
	return &Password{
		Value: value,
	}, nil
}
