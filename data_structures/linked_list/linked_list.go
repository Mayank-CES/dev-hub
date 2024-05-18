package main

import "fmt"

// node
type node struct {
	data int
	next *node
}

// linkedList
type linkedList struct {
	head   *node
	length int
}

// prepend
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// deleteWithValue
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head
	for prevToDelete.next.data != value {
		if prevToDelete.next.next == nil {
			return
		}
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}

// printListData
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	list := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 4}
	node5 := &node{data: 5}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)
	fmt.Println(list)    // output:  {0xc000104060 5}
	list.printListData() // output:  5 4 3 2 1
	list.deleteWithValue(3)
	list.printListData() // output:  5 4 2 1
}
