package main

import (
	"errors"
	"regexp"
)

type CPF struct {
	value string
}

func (e *CPF) NewCPF(value string) (*CPF, error) {
	matchEmail := regexp.MustCompile("^(.+)@(.+)$").MatchString(value)
	if !matchEmail {
		return nil, errors.New("invalid email")
	}
	return &CPF{
		value: value,
	}, nil
}
