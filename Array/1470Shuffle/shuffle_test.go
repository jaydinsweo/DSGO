package main

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {

	got := Shuffle([]int{1, 2, 3, 4, 1, 2, 3, 4}, 4)
	want := []int{1, 1, 2, 2, 3, 3, 4, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
