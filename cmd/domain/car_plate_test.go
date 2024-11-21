package domain_test

import (
	"testing"

	"github.com.br/gibranct/ride/cmd/domain"
	"github.com/stretchr/testify/assert"
)

func Test_CreateValidCarPlate(t *testing.T) {
	carPlates := []string{"ABC9090", "AAA1111"}
	for _, n := range carPlates {
		carPlate, err := domain.NewCarPlate(n)
		assert.Nil(t, err)
		assert.Equal(t, n, carPlate.GetValue())
	}
}

func Test_CreateInvalidCarPlate(t *testing.T) {
	invalidCarPlates := []string{"ABC909", "AA1111", "A1A1111", "AAA11B1"}
	for _, n := range invalidCarPlates {
		carPlate, err := domain.NewCarPlate(n)
		assert.NotNil(t, err)
		assert.Nil(t, carPlate)
	}
}
