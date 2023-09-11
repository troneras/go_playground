package main

import (
	"testing"
)

func TestAppend(t *testing.T) {
	list := makeList([]int{1, 2, 3, 4, 5})

	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, got %d", list.Size())
	}

	list.Append(6)
	list.Append(7)

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	actual := listToSlice(list)
	if !sliceEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	if list.Size() != 7 {
		t.Errorf("Expected list size to be 7, got %d", list.Size())
	}
}

func TestInsert(t *testing.T) {
	list := makeList([]int{1, 2, 3, 4, 5})

	list.Insert(100, 0)
	list.Insert(200, 2)
	list.Insert(300, 5)
	list.Insert(400, 7)

	if list.Size() != 9 {
		t.Errorf("Expected list size to be 9, got %d", list.Size())
	}

	err := list.Insert(500, -1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	position, _ := list.Search(100)
	if position != 0 {
		t.Errorf("Expected position to be 0, got %d", position)
	}

	position, _ = list.Search(200)
	if position != 2 {
		t.Errorf("Expected position to be 2, got %d", position)
	}

	expected := []int{100, 1, 200, 2, 3, 300, 4, 400, 5}
	actual := listToSlice(list)
	if !sliceEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func TestSearch(t *testing.T) {
	list := makeList([]int{1, 2, 3, 4, 5})

	position, _ := list.Search(1)
	if position != 0 {
		t.Errorf("Expected position to be 0, got %d", position)
	}

	position, _ = list.Search(3)
	if position != 2 {
		t.Errorf("Expected position to be 2, got %d", position)
	}

	position, _ = list.Search(5)
	if position != 4 {
		t.Errorf("Expected position to be 4, got %d", position)
	}

	_, err := list.Search(6)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestDelete(t *testing.T) {
	list := makeList([]int{1, 2, 3, 4, 5})

	list.Delete(0)

	if list.Size() != 4 {
		t.Errorf("Expected list size to be 4, got %d", list.Size())
	}

	list.Delete(2)

	if list.Size() != 3 {
		t.Errorf("Expected list size to be 3, got %d", list.Size())
	}

	err := list.Delete(3)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	err = list.Delete(-1)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	expected := []int{2, 3, 5}
	actual := listToSlice(list)
	if !sliceEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
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
