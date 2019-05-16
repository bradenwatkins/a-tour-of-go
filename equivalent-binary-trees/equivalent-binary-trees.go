package main

import (
	"golang.org/x/tour/tree"
)

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool { 
	ch1 := make(chan int)
	ch2 := make(chan int)
	vals := make(map[int]int)
	go Walker(t1, ch1)
	go Walker(t2, ch2)
	for v := range ch1 {
		vals[v] += 1
	}
	for v := range ch2 {
		if vals[v] == 0 {
			return false
		} else {
			vals[v] -= 1
		}
	}
	return true
}

func main() {
	println(Same(tree.New(1), tree.New(1)))
	println(Same(tree.New(1), tree.New(2)))	
}
