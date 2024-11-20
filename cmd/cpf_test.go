package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CPFWithDigitDifferentThanZero(t *testing.T) {
	cpf, err := NewCPF("97456321558")
	assert.Nil(t, err)
	assert.NotEmpty(t, cpf.value)
}

func Test_CPFWithZeroDigit(t *testing.T) {
	cpf, err := NewCPF("71428793860")
	assert.Nil(t, err)
	assert.NotEmpty(t, cpf.value)
}

func Test_CPFWithZeroAsFirstDigit(t *testing.T) {
	cpf, err := NewCPF("87748248800")
	assert.Nil(t, err)
	assert.NotEmpty(t, cpf.value)
}

func Test_CPFWithLessThan11Digits(t *testing.T) {
	cpf, err := NewCPF("9745632155")
	assert.NotNil(t, err)
	assert.Nil(t, cpf)
}

func Test_CPFWithDigitsTheSame(t *testing.T) {
	cpf, err := NewCPF("11111111111")
	assert.NotNil(t, err)
	assert.Nil(t, cpf)
}

func Test_CPFWithLetters(t *testing.T) {
	cpf, err := NewCPF("97a56321558")
	assert.NotNil(t, err)
	assert.Nil(t, cpf)
}
