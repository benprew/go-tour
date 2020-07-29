package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkIter(t, ch)
	close(ch)
}

func walkIter(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkIter(t.Left, ch)
	ch <- t.Value
	walkIter(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		j := <-ch2
		if i != j {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)

	fmt.Printf("%v\n%v\n", t1, t2)
	fmt.Println(Same(t1, t2))
}
