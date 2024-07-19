package main

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{3, 1, 5, 2, 5, 6, 7, 8, 3, 1, 1, 6, 7, 8, 23, 53, 65, 32}
	want := append([]int{}, arr...)
	slices.Sort(want)
	
	got := selectionSort(arr)

	if !slices.Equal(got, want) {
		t.Fatalf("got %v\nwant %v", got, want)
	}
}
