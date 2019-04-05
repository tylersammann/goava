package sets

import (
	"reflect"
	"testing"

	"fmt"
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

func TestCopy_emptySet(t *testing.T) {
	expected := New()
	actual := expected.Copy()

	assertMatchingValues(t, expected, actual)
	assert.False(t, expected == actual)
}

func TestCopy(t *testing.T) {
	expected := New(1.2, 3.4)
	actual := expected.Copy()

	assertMatchingValues(t, expected, actual)
	assert.False(t, expected == actual)
}

func TestForEach_emptySet(t *testing.T) {
	actual := New()

	actual.ForEach(func(item interface{}) {
		assert.Fail(t, "The empty set should have no items. The method should not get called")
	})
}

func TestForEach(t *testing.T) {
	numbers := New(8, 9, 10, 11)
	expected := New(18, 19, 20, 21)
	actual := New()

	numbers.ForEach(func(item interface{}) {
		intItem := item.(int)
		actual.Add(intItem + 10)
	})

	assertMatchingValues(t, expected, actual)
}

func TestFindFirst_emptySet(t *testing.T) {
	actual := New()

	actual.FindFirst(func(item interface{}) bool {
		assert.Fail(t, "The empty set should have no items. The method should not get called")
		return false
	})
}

func TestFindFirst_everyItemMeetsCondition(t *testing.T) {
	numbers := New(8, 9)
	calls := 0

	actual := numbers.FindFirst(func(item interface{}) bool {
		calls++
		return item.(int) > 1
	})

	assert.Equal(t, 1, calls)
	assert.True(t, actual == 8 || actual == 9)
}

func TestFindFirst(t *testing.T) {
	numbers := New(8, 9, 10, 11)
	calls := 0

	actual := numbers.FindFirst(func(item interface{}) bool {
		calls++
		return item.(int) > 9
	})

	assert.True(t, calls > 0 && calls <= 3)
	assert.True(t, actual == 10 || actual == 11)
}

func assertMatchingValues(t *testing.T, expected Set, actual Set) {
	assert.Equal(t, expected.Size(), actual.Size(), "expected size should equal actual size")
	assert.Equal(t, expected.rType(), actual.rType(), "expected type should equal actual type")
	expected.ForEach(func(item interface{}) {
		assert.True(t, actual.Has(item), fmt.Sprintf("actual %v should contain expected %v", actual, item))
	})
}
