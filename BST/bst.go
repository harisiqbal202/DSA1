package main

import "fmt"

var count int

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(k int) {
	if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	} else if n.key < k {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	}
}
func (n *Node) Search(k int) bool {
	count++

	if n == nil {
		return false
	}

	if n.key < k {
		return n.right.Search(k)
	} else if n.key > k {
		return n.left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{key: 100}
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(200)
	tree.Insert(300)
	fmt.Println(tree)

	fmt.Println(count)
	x := tree.Search(400)
	fmt.Println(x)
}
