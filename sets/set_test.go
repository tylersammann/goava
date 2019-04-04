package sets

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
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

func TestAdd_oneItem(t *testing.T) {
	actual := New("this", "was")
	actual.Add("added")

	assertMatchingValues(t, New("this", "was", "added"), actual)
}

func TestAdd_multipleItems_duplicateItems_zeroItems(t *testing.T) {
	actual := New("this", "was")
	actual.Add("also", "added").Add("too")
	actual.Add("also", "this").Add()

	assertMatchingValues(t, New("this", "was", "also", "added", "too"), actual)
}

func TestAdd_wrongTypeItems(t *testing.T) {
	actual := New("this", "was")

	assert.Panics(t, func() {
		actual.Add(1)
	})

	assert.Panics(t, func() {
		actual.Add(2.3)
	})

	assert.Panics(t, func() {
		actual.Add(true)
	})

	assert.Panics(t, func() {
		actual.Add("hey", 2, 5.3)
	})
}

func TestRemove_oneItem(t *testing.T) {
	actual := New(1, 2, 3)
	actual.Remove(2)

	assertMatchingValues(t, New(1, 3), actual)
}

func TestRemove_multipleItems_duplicateItems_zeroItems(t *testing.T) {
	actual := New(1, 2, 3)
	actual.Remove(1, 2).Remove(2).Remove()

	assertMatchingValues(t, New(3), actual)
}

func TestRemove_fromEmptySet(t *testing.T) {
	actual := New(1, 2, 3)
	actual.Remove(1, 2, 3)

	assertMatchingValues(t, New(), actual)

	actual.Remove(3)
	assertMatchingValues(t, New(), actual)
}

func assertMatchingValues(t *testing.T, expected Set, actual Set) {
	assert.Equal(t, expected.Size(), actual.Size(), "expected size should equal actual size")
	assert.Equal(t, expected.rType(), actual.rType(), "expected type should equal actual type")
	expected.ForEach(func(item interface{}) {
		assert.True(t, actual.Has(item), fmt.Sprintf("actual %v should contain expected %v", actual, item))
	})
}
