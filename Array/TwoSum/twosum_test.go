package main

import (
	"reflect"
	"testing"
)

// Given an array []int{} and a target int
// returns the indices of the two numbers that add up to the target
// Assume: exactly one solution available

func TestTwoSum(t *testing.T) {

	checkEq := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("same array", func(t *testing.T) {

		got := TwoSum([]int{2, 4, 7, 5, 9}, 13)
		want := []int{1, 4}

		checkEq(t, got, want)
	})

}
