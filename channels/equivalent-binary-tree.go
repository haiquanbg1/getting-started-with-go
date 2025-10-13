package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var next func(t *tree.Tree)

	next = func(tr *tree.Tree) {
		if tr == nil {
			return
		}

		next(tr.Left)
		ch <- tr.Value
		next(tr.Right)
	}

	next(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, end1 := <-ch1
		v2, end2 := <-ch2

		if end1 != end2 {
			return false
		}

		if !end1 || !end2 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
