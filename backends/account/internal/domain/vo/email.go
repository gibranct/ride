package vo

import (
	"regexp"

	"github.com.br/gibranct/account/internal/domain/errors"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	matchEmail := regexp.MustCompile("^(.+)@(.+)$").MatchString(value)
	if !matchEmail {
		return nil, errors.ErrInvalidEmail
	}
	return &Email{
		value: value,
	}, nil
}

func (e *Email) GetValue() string {
	return e.value
}
