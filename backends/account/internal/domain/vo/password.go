package vo

import (
	"unicode/utf8"

	"github.com.br/gibranct/account/internal/domain/errors"
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
		return nil, errors.ErrPasswordTooShort
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Password{
		Value: string(hashedPassword),
	}, nil
}
