package list

import (
	"testing"
)

func TestAppend(t *testing.T) {

	tests := []struct {
		name     string
		initial  []int
		append   []int
		expected []int
	}{
		{
			name:     "Append 6 and 7",
			initial:  []int{1, 2, 3, 4, 5},
			append:   []int{6, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel() // marks the test case as capable of running in parallel

			list := MakeList(test.initial)
			initialSize := len(test.initial)

			if list.Size() != initialSize {
				t.Errorf("Expected list size to be 5, got %d", list.Size())
			}

			for _, v := range test.append {
				list.Append(v)
			}

			if list.Size() != len(test.expected) {
				t.Errorf("Expected list size to be 7, got %d", list.Size())
			}
			actual := listToSlice(list)

			if !sliceEqual(test.expected, actual) {
				t.Errorf("Expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		insert  []struct {
			value    int
			position int
		}
		search []struct {
			value    int
			position int
		}
		expected []int
	}{
		{
			name:     "Insert 100, 200, 300, 400",
			initial:  []int{1, 2, 3, 4, 5},
			insert:   []struct{ value, position int }{{100, 0}, {200, 2}, {300, 5}, {400, 7}},
			search:   []struct{ value, position int }{{100, 0}, {200, 2}, {300, 5}, {400, 7}},
			expected: []int{100, 1, 200, 2, 3, 300, 4, 400, 5},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			list := MakeList(v.initial)

			for _, v := range v.insert {
				list.Insert(v.value, v.position)
			}

			if list.Size() != len(v.expected) {
				t.Errorf("Expected list size to be 9, got %d", list.Size())
			}

			for _, v := range v.search {
				position, _ := list.Search(v.value)
				if position != v.position {
					t.Errorf("Expected position to be %d, got %d", v.position, position)
				}
			}

			actual := listToSlice(list)
			if !sliceEqual(v.expected, actual) {
				t.Errorf("Expected %v, got %v", v.expected, actual)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name    string
		initial []int
		search  []struct {
			value    int
			position int
		}
	}{
		{
			name:    "Search 1, 2, 3, 4, 5",
			initial: []int{1, 2, 3, 4, 5},
			search: []struct{ value, position int }{
				{1, 0},
				{2, 1},
				{3, 2},
				{4, 3},
				{5, 4},
			},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			list := MakeList(v.initial)

			for _, v := range v.search {
				position, _ := list.Search(v.value)
				if position != v.position {
					t.Errorf("Expected position to be %d, got %d", v.position, position)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		delete   []int
		expected []int
	}{
		{
			name:     "Delete 0, 2, 3",
			initial:  []int{1, 2, 3, 4, 5},
			delete:   []int{0, 2, 3},
			expected: []int{2, 3, 5},
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			list := MakeList(v.initial)

			for _, v := range v.delete {
				list.Delete(v)
			}

			if list.Size() != len(v.expected) {
				t.Errorf("Expected list size to be 3, got %d", list.Size())
			}

			actual := listToSlice(list)
			if !sliceEqual(v.expected, actual) {
				t.Errorf("Expected %v, got %v", v.expected, actual)
			}
		})
	}
}

func listToSlice[T comparable](list *List[T]) []T {
	var result []T
	node := list.head
	for node != nil {
		result = append(result, node.val)
		node = node.next
	}
	return result
}

func sliceEqual[T comparable](a, b []T) bool {
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

/**
 * BENCHMARKS
 * runs with: go test -bench
 */
func BenchmarkAppend(b *testing.B) {
	list := MakeList([]int{})
	for i := 0; i < b.N; i++ {
		list.Append(i)
	}
}
