package vo

import (
	"errors"
	"regexp"
)

type CarPlate struct {
	value string
}

func NewCarPlate(value string) (*CarPlate, error) {
	matchCarPlate := regexp.MustCompile("[A-Z]{3}[0-9]{4}").MatchString(value)
	if !matchCarPlate {
		return nil, errors.New("invalid car plate")
	}
	return &CarPlate{
		value: value,
	}, nil
}

func (cp *CarPlate) GetValue() string {
	return cp.value
}
