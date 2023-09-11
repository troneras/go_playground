package list

type SortableList[T comparable] interface {
	ListOps[T]
	Sort()
}

type SortedList[T comparable] struct {
	List[T]
}

type Comparer[T any] func(a, b T) int

func (l *SortedList[T]) Sort(cmp Comparer[T]) {
	if l.Size() <= 1 {
		return
	}

	var sorted bool
	for !sorted {
		sorted = true
		for i := 0; i < l.Size()-1; i++ {
			if cmp(l.Get(i).val, l.Get(i+1).val) > 0 {
				l.Swap(i, i+1)
				sorted = false
			}
		}
	}
}

func MakeSortedList[T comparable](elements []T) *SortedList[T] {
	l := &SortedList[T]{}
	for _, v := range elements {
		l.Append(v)
	}

	return l
}
