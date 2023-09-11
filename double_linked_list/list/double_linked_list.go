package list

import (
	"fmt"
)

type Node[T comparable] struct {
	next *Node[T]
	prev *Node[T]
	val  T
}

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type ListOps[T comparable] interface {
	Insert(element T, position int) *MyError
	Delete(position int) *MyError
	Search(element T) (int, *MyError)
	Size() int
	Append(v T)
}

type MyError struct {
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("An error ecurred: %s", e.What)
}

func (l *Node[T]) getNext() (T, *MyError) {
	if l.next != nil {
		return l.next.val, nil
	}
	var zero T
	return zero, &MyError{"No next element"}
}

func (l *Node[T]) hasNext() bool {
	return l != nil && l.next != nil
}

func (l *Node[T]) hasPrev() bool {
	return l != nil && l.prev != nil
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Insert(element T, position int) *MyError {
	if position < 0 || position > l.Size() {
		return &MyError{"Position out of range"}
	}

	if position == l.Size() {
		fmt.Println("Appending")
		l.Append(element)
		return nil
	}

	newNode := &Node[T]{val: element}
	if position == 0 {
		// Update head and possibly tail
		newNode.next = l.head
		if l.head != nil {
			l.head.prev = newNode
		}
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
	} else {
		node := l.head
		for pos := 0; pos < position; pos++ {
			node = node.next
		}

		if node.next != nil {
			node.prev.next = newNode
			newNode.prev = node.prev
			node.prev = newNode
			newNode.next = node
		} else {
			newNode.prev = node.prev
			newNode.next = node
			node.prev.next = newNode
			node.prev = newNode
		}
	}
	l.size++
	return nil
}

func (l *List[T]) Delete(position int) *MyError {
	if position < 0 || position >= l.Size() {
		return &MyError{"Position out of range"}
	}

	if position == 0 { // update head and possibly tail
		if l.head.next != nil {
			l.head = l.head.next
			l.head.prev = nil
			l.size--
		} else {
			l.head = nil
			l.tail = nil
			l.size = 0
		}

		return nil
	}
	if position == l.Size() {
		l.tail = l.tail.prev
		l.size--
		return nil
	}

	// normal delete (I could do front or back delete, but for simplicity...)
	node := l.head
	for pos := 0; pos < position; pos++ {
		node = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
		node.prev.next = node.next
	} else {
		node.prev.next = nil
	}

	l.size--

	return nil
}

func (l *List[T]) Search(element T) (int, *MyError) {
	node := l.head
	for i := 0; i < l.Size(); i++ {
		if node.val == element {
			return i, nil
		}
		node = node.next
	}

	var zero int
	return zero, &MyError{"Element not found"}
}

func PrintList[T comparable](list *List[T]) {
	if list == nil {
		return
	}
	fmt.Println("List size : ", list.Size())

	node := list.head // head node
	for i := 0; i < list.Size(); i++ {
		fmt.Printf("Node %d is %v\n", i, node.val)
		node = node.next
	}
}

func PrintReverse[T comparable](list *List[T]) {
	if list == nil {
		return
	}
	PrintListSize(list)

	node := list.tail // head node

	for i := list.Size(); i > 0; i-- {
		fmt.Printf("Node %d is %v\n", i, node.val)
		node = node.prev
	}
}

func PrintListSize[T comparable](l ListOps[T]) {
	fmt.Println("List size:", l.Size())
}

func (l *List[T]) Append(v T) {
	if l == nil {
		return
	}
	var node = &Node[T]{val: v, prev: l.tail}
	if l.head == nil {
		l.head = node
	}
	if l.tail != nil {
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node

	l.size++
}

func MakeList[T comparable](values []T) *List[T] {
	list := &List[T]{}
	for _, v := range values {
		list.Append(v)
	}
	return list
}
