package vo

import (
	"errors"
	"regexp"
)

type Name struct {
	value string
}

func NewName(value string) (*Name, error) {
	matchName := regexp.MustCompile("[a-zA-Z] [a-zA-Z]+").MatchString(value)
	if !matchName {
		return nil, errors.New("invalid name")
	}
	return &Name{
		value: value,
	}, nil
}

func (n *Name) GetValue() string {
	return n.value
}
