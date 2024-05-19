package main

import "fmt"

// Stack represents a stack that hold a slice
type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and return the removed value
func (s *Stack) Pop() int {
	l := len(s.items)
	removedItem := s.items[l-1]
	s.items = s.items[:l-1]
	return removedItem
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack)       // Output : 	{[1 2 3 4]}
	fmt.Println(stack.Pop()) // Output : 	4
	fmt.Println(stack)       // Output : 	{[1 2 3]}

}
