//test sort.go
package main

import "testing"
import "sort"

func TestMergeSort(t *testing.T) {

	numbers := []int{5, 6, 1, 4, 9, 8, 8, 2, 7}
	result := mergeIntArraySort(numbers)
	sort.Ints(numbers)

	if !Equal(numbers, result) {
		t.Errorf("expected '%v' but got '%v'", numbers, result)
	}

}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}