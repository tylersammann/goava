package sets

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_withEmtpySet(t *testing.T) {
	actual := New()
	assert.Equal(t, 0, actual.Size())
}

func TestNew_withIntSet_oneItem(t *testing.T) {
	actual := New(1)
	expectedValues := []int{1}
	assert.Equal(t, reflect.TypeOf(0), actual.rType())
	assert.Equal(t, len(expectedValues), actual.Size())

	for _, item := range expectedValues {
		assert.True(t, actual.Has(item))
	}
}

func TestNew_withIntSet_manyItems(t *testing.T) {
	actual := New(1, 5, 7, 34)
	expectedValues := []int{1, 5, 7, 34}
	assert.Equal(t, reflect.TypeOf(0), actual.rType())
	assert.Equal(t, len(expectedValues), actual.Size())

	for _, item := range expectedValues {
		assert.True(t, actual.Has(item))
	}
}

func TestNew_withStringSet_manyItems(t *testing.T) {
	actual := New("this", "is", "reaLLY", "cool")
	expectedValues := []string{"this", "is", "reaLLY", "cool"}
	assert.Equal(t, reflect.TypeOf(""), actual.rType())
	assert.Equal(t, len(expectedValues), actual.Size())

	for _, item := range expectedValues {
		assert.True(t, actual.Has(item))
	}
}

func TestNew_withItemsOfDifferentTypes(t *testing.T) {
	assert.Panics(t, func() {
		New("this", 234, "is", "bad")
	})
}
