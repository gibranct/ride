package di_test

import (
	"testing"

	di "github.com.br/gibranct/ride/cmd/infra/DI"
	"github.com/stretchr/testify/assert"
)

func Test_Registry(t *testing.T) {
	registry := di.NewRegistry()

	if assert.NotNil(t, registry) {
		registry.Add("a", "1")
		registry.Add("b", "2")
		registry.Add("c", "3")

		obj, err := registry.Get("a")
		assert.Nil(t, err)
		assert.Equal(t, "1", obj)

		obj, err = registry.Get("b")
		assert.Nil(t, err)
		assert.Equal(t, "2", obj)

		obj, err = registry.Get("c")
		assert.Nil(t, err)
		assert.Equal(t, "3", obj)

		registry.Remove("a")
		registry.Remove("b")
		registry.Remove("c")

		obj, err = registry.Get("a")
		assert.Equal(t, "object was not registered for key: a", err.Error())
		assert.Nil(t, obj)

		obj, err = registry.Get("b")
		assert.Equal(t, "object was not registered for key: b", err.Error())
		assert.Nil(t, obj)

		obj, err = registry.Get("c")
		assert.Equal(t, "object was not registered for key: c", err.Error())
		assert.Nil(t, obj)
	}
}
