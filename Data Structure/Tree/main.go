package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func main() {
	tree := &Node{Key: 100}

	tree.Insert(123)
	tree.Insert(23)
	tree.Insert(400)
	tree.Insert(124)
	tree.Insert(24)
	tree.Insert(500)
	fmt.Println(tree)
	fmt.Println(tree.Search(23))
	fmt.Println(count)
}

func (n *Node) Insert(k int) {
	if k > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if k < n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {

	count++

	if n == nil {
		return false
	}
	if k > n.Key {
		return n.Right.Search(k)
	} else if k < n.Key {
		return n.Left.Search(k)
	}
	return true
}
