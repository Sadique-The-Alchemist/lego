package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		got := Sum(values)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, values)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sum(numbers)
		want := 28
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 3}, []int{4, 5})
	want := []int{5, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("test sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{4, 5})
		want := []int{3, 5}
		checkSums(t, got, want)
	})
	t.Run("test the sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5})
		want := []int{0, 5}
		checkSums(t, got, want)
	})

}
