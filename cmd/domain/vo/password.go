package vo

import (
	"errors"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Value string
}

func (p *Password) Compare(hashedPassword, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(value))
	return err == nil
}

func NewPassword(value string) (*Password, error) {
	if utf8.RuneCountInString(value) < 6 {
		return nil, errors.New("invalid password: must be greater than 5")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Password{
		Value: string(hashedPassword),
	}, nil
}
