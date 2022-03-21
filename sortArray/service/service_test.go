package service

import "testing"

func TestShuffleSort(t *testing.T) {
	unsortedSlice := []int{3, 1, 4}
	sortedSlice := []int{1, 3, 4}
	ShuffleSort(unsortedSlice)
	if !checkSlices(unsortedSlice, sortedSlice) {
		t.Fatal("the slice has not been correctly sorted")
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
