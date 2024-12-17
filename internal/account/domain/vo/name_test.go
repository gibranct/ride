package vo_test

import (
	"testing"

	"github.com.br/gibranct/ride/internal/account/domain/vo"
	"github.com/stretchr/testify/assert"
)

func Test_CreateValidName(t *testing.T) {
	validNames := []string{"John Doe", "Gil Bil"}
	for _, n := range validNames {
		name, err := vo.NewName(n)
		assert.Nil(t, err)
		assert.Equal(t, n, name.GetValue())
	}
}

func Test_CreateInvalidName(t *testing.T) {
	invalidNames := []string{"John", "Gil", "", "Fsads "}
	for _, n := range invalidNames {
		name, err := vo.NewName(n)
		assert.NotNil(t, err)
		assert.Nil(t, name)
	}
}
