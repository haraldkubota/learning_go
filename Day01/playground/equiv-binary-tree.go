package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Traverse(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Traverse(t1, c1)
	go Traverse(t2, c2)

	res1 := ""
	res2 := ""
	for {
		select {
		case v1, ok := <-c1:
			if !ok {
				c1 = nil
			} else {
				res1 = res1 + " " + fmt.Sprintf("%d", v1)
			}
		case v2, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				res2 = res2 + " " + fmt.Sprintf("%d", v2)
			}
		}
		if c1 == nil && c2 == nil {
			break
		}
	}
	fmt.Printf("res1=%s, res2=%s", res1, res2)
	return res1 == res2
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)
	fmt.Println("Same: ", Same(t1, t2))
	fmt.Println("Same: ", Same(t1, t3))
}
