package main

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func (e *Email) NewEmail(value string) (*Email, error) {
	matchEmail := regexp.MustCompile("^(.+)@(.+)$").MatchString(value)
	if !matchEmail {
		return nil, errors.New("invalid email")
	}
	return &Email{
		value: value,
	}, nil
}
