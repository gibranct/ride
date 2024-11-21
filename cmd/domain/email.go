package domain

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	matchEmail := regexp.MustCompile("^(.+)@(.+)$").MatchString(value)
	if !matchEmail {
		return nil, errors.New("invalid email")
	}
	return &Email{
		value: value,
	}, nil
}

func (e *Email) GetValue() string {
	return e.value
}
