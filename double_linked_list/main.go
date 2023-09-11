package main

import (
	"fmt"
	"go_playground/list"
)

func main() {
	l := list.MakeList([]int{1, 2, 3})
	var err *list.MyError

	// Print list (0, 1, 2, 3)
	list.PrintList(l)

	var inserts = [][2]int{{5, 2}, {4, 4}, {4, 6}, {3, 0}, {4, -1}}

	for _, v := range inserts {
		// Insert function
		err = l.Insert(v[0], v[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Print list
		list.PrintList(l)
	}

	var deletes = []int{0, 1, 2}
	for _, v := range deletes {
		// Insert function
		err = l.Delete(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Print list
		list.PrintList(l)
	}

	// search method
	var pos int
	pos, err = l.Search(5)
	if pos, err = l.Search(5); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Element 5 is at position ", pos)
	}

	// reverse method
	list.PrintReverse(l)
}
