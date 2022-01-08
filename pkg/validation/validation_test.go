package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsNil(t *testing.T) {
	t.Run("should return an error if pointer is nil", func(t *testing.T) {
		type randomType struct{}
		var value *randomType
		err := IsRequired(value, "some element")
		require.Error(t, err)
		assert.EqualError(t, err, "some element is nil")
	})

	t.Run("should return an error if string is empty", func(t *testing.T) {
		err := IsRequired("", "some element")
		require.Error(t, err)
		assert.EqualError(t, err, "some element is empty")
	})

	t.Run("should return an error if slice is empty", func(t *testing.T) {
		err := IsRequired([]string{}, "some element")
		require.Error(t, err)
		assert.EqualError(t, err, "some element is empty")
	})

	t.Run("should return an error if map is empty", func(t *testing.T) {
		someMap := map[string]bool{}
		err := IsRequired(someMap, "some element")
		require.Error(t, err)
		assert.EqualError(t, err, "some element is empty")
	})

	t.Run("should return an error if map is nil", func(t *testing.T) {
		var someMap map[string]bool
		err := IsRequired(someMap, "some element")
		require.Error(t, err)
		assert.EqualError(t, err, "some element is empty")
	})

	t.Run("should return no error if element is not nil", func(t *testing.T) {
		err := IsRequired("some string", "some element")
		assert.NoError(t, err)
	})

	t.Run("should return error if element is an interface and it's nil", func(t *testing.T) {
		type someInterface interface {
			Do() error
		}
		var value someInterface
		err := IsRequired(value, "some element")
		require.Error(t, err)
		assert.EqualError(t, err, "some element is nil")
	})
}

func sPtr(s string) *string { return &s }

func TestHasNoNilElements(t *testing.T) {
	t.Run("should return no error if the target is not a slice or array", func(t *testing.T) {
		assert.Nil(t, HasNoNilElements(""))
	})

	t.Run("should return no error if the target is a container with no elements", func(t *testing.T) {
		assert.Nil(t, HasNoNilElements([]string{}))
	})
	t.Run("should return no error if the target is a container and none element is nil", func(t *testing.T) {
		assert.Nil(t, HasNoNilElements([]*string{sPtr("123"), sPtr("321")}))
	})

	t.Run("should return errors for nil elements of a valid container", func(t *testing.T) {
		container := []*string{nil, nil, sPtr("123")}
		errs := HasNoNilElements(container)
		assert.NotNil(t, errs)
		assert.EqualError(t, errs, "2 errors:[ index 0 is nil, index 1 is nil ]")
	})
}
