package validator_test

import (
	"testing"

	"github.com.br/gibranct/ride/pkg/validator"
	"github.com/stretchr/testify/assert"
)

func Test_CPFWithDigitDifferentThanZero(t *testing.T) {
	cpf := "97456321558"
	isValid := validator.ValidateCPF(cpf)
	assert.True(t, isValid)
}

func Test_CPFWithZeroDigit(t *testing.T) {
	cpf := "71428793860"
	isValid := validator.ValidateCPF(cpf)
	assert.True(t, isValid)
}

func Test_CPFWithZeroAsFirstDigit(t *testing.T) {
	cpf := "87748248800"
	isValid := validator.ValidateCPF(cpf)
	assert.True(t, isValid)
}

func Test_CPFWithLessThan11Digits(t *testing.T) {
	cpf := "9745632155"
	isValid := validator.ValidateCPF(cpf)
	assert.False(t, isValid)
}

func Test_CPFWithDigitsTheSame(t *testing.T) {
	cpf := "11111111111"
	isValid := validator.ValidateCPF(cpf)
	assert.False(t, isValid)
}

func Test_CPFWithLetters(t *testing.T) {
	cpf := "97a56321558"
	isValid := validator.ValidateCPF(cpf)
	assert.False(t, isValid)
}
