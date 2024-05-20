package main

import "fmt"

// Node represents the components of a Binary Search Tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and Return true if there is a node with that value
func (n *Node) Search(k int) bool {
	// if node is empty we traversed the whole tree
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return n.Key == k
}

func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree)
	// tree.Insert(50)
	// fmt.Println(tree)
	tree.Insert(200)
	fmt.Println(tree)
	tree.Insert(130)
	fmt.Println(tree)
	/*
		&{100 <nil> <nil>}
		&{100 0xc000010048 <nil>}
		&{100 0xc000010048 0xc000010078}
	*/
	tree.Insert(2)
	tree.Insert(230)
	tree.Insert(53)
	tree.Insert(12)
	tree.Insert(349)
	tree.Insert(394)
	tree.Insert(293)
	tree.Insert(193)

	fmt.Println(tree.Search(293)) // output : true
	fmt.Println(tree.Search(223)) // output : false

}
