package service

import (
	"fmt"
	"testing"
)

func TestShuffleSort_shouldSortSlice_whenNormalSliceProvided(t *testing.T) {
	unsortedSlice := []int{3, 1, 4}
	sortedSlice := []int{1, 3, 4}
	ShuffleSort(unsortedSlice)
	if !checkSlices(unsortedSlice, sortedSlice) {
		t.Fatal(fmt.Sprintf("The slice has not been correctly sorted!\nExpected: %v\nGot:%v", sortedSlice, unsortedSlice))
	}
}

func TestShuffleSort_emptySlice_isSorted(t *testing.T) {
	unsortedSlice := []int{}
	sortedSlice := []int{}
	ShuffleSort(unsortedSlice)
	if !checkSlices(unsortedSlice, sortedSlice) {
		t.Fatal(fmt.Sprintf("The slice should be empty and therefore already sorted!\nExpected: %v\nGot:%v", sortedSlice, unsortedSlice))
	}
}

func checkSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
