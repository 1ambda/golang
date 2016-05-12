package concurrencies

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/*
  type Tree struct {
  	Left  *Tree
  	Value int
  	Right *Tree
  }

  tree.New(k) constructs randomly-structured binary tree holding the values k, ..., 10k
*/

func _walk(t *tree.Tree, ch chan int) {
	if t != nil { // in-order
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func ExampleWalk1() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if (!ok1 || !ok2) || (v1 != v2) {
			return false
		}

		if ok1 && ok2 {
			return true
		}
	}
}

func ExampleSame1() {
	b1 := Same(tree.New(1), tree.New(1))
	b2 := Same(tree.New(1), tree.New(2))
	b3 := Same(tree.New(5), tree.New(3))

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

	// Output:
	// true
	// false
	// false
}
