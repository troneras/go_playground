package list

import (
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		expected []int
	}{
		{
			name:     "Sort 3, 4, 1, 2, 5",
			initial:  []int{3, 4, 1, 2, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, v := range tests {
		v := v
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			list := MakeSortedList(v.initial)
			list.Sort(func(a, b int) int {
				return a - b
			})

			if !SliceEqual(ListToSlice(list), v.expected) {
				t.Errorf("Expected %v, got %v", v.expected, ListToSlice(list))
			}
		})
	}
}
