package main

import "fmt"

// Queue represents a queue that hold a slice
type Queue struct {
	items []int
}

// Enqueue will add a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue will remove a value at the front
// and return the removed value
func (q *Queue) Dequeue() int {
	removedItem := q.items[0]
	q.items = q.items[1:len(q.items)]
	return removedItem
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue)           // Output : 	{[1 2 3 4]}
	fmt.Println(queue.Dequeue()) // Output : 	1
	fmt.Println(queue)           // Output : 	{[2 3 4]}

}
