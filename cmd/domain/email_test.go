package domain_test

import (
	"testing"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com/stretchr/testify/assert"
)

func Test_CreateValidEmail(t *testing.T) {
	validEmails := []string{"john@doe.com", "gil@bil.com"}
	for _, n := range validEmails {
		email, err := domain.NewEmail(n)
		assert.Nil(t, err)
		assert.Equal(t, n, email.GetValue())
	}
}

func Test_CreateInvalidEmails(t *testing.T) {
	invalidEmails := []string{"johndoecom", "gilbil.com"}
	for _, n := range invalidEmails {
		email, err := domain.NewEmail(n)
		assert.NotNil(t, err)
		assert.Nil(t, email)
	}
}
