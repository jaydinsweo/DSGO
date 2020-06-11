package numbers

import "testing"

// given an array nums of integers
// return integers of how many of them contain an even number of digits

func TestFindNumbers(t *testing.T) {

	t.Run("returns 3 answers", func(t *testing.T) {
		got := FindNumbers([]int{123, 4, 53, 64, 4354})
		want := 3

		if got != want {
			t.Errorf("got %v want %d \n", got, want)
		}
	})

	t.Run("return none", func(t *testing.T) {
		got := FindNumbers([]int{123, 123, 123, 1, 12332})
		want := 0

		if got != want {
			t.Errorf("Return none \n got %v want %d \n", got, want)
		}

	})
}
