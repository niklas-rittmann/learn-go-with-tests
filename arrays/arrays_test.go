package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := 15
	got := Sum(arr)
	if got != expected {
		t.Errorf("Got %d expected %d given %v", got, expected, arr)
	}

}

func TestSumAll(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2}
	expected := []int{15, 3}
	got := SumAll(arr1, arr2)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got %d expected %d given %v and %v", got, expected, arr1, arr2)
	}

}

func TestSumAllTails(t *testing.T) {
	assertValidateInput := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %d expected %d", got, expected)
		}
	}

	t.Run("Run with valid input", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := []int{1, 2}
		expected := []int{14, 2}
		got := SumAllTails(arr1, arr2)
		assertValidateInput(t, got, expected)

	})
	t.Run("Run with empty slice", func(t *testing.T) {
		arr1 := []int{}
		arr2 := []int{1, 2}
		expected := []int{0, 2}
		got := SumAllTails(arr1, arr2)
		assertValidateInput(t, got, expected)

	})

}
