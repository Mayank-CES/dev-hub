package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array is empty")
		return -1
	}

	// take out the last index and put it in the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[parent(index)] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m) // output :		&{[]}
	buildHeap := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

	/* Full Logs:
	&{[]}
	&{[1]}
	&{[2 1]}
	&{[3 1 2]}
	&{[4 3 2 1]}
	&{[5 4 2 1 3]}
	&{[6 4 5 1 3 2]}
	&{[7 4 6 1 3 2 5]}
	&{[8 7 6 4 3 2 5 1]}
	&{[9 8 6 7 3 2 5 1 4]}
	&{[10 9 6 7 8 2 5 1 4 3]}
	&{[11 10 6 7 9 2 5 1 4 3 8]}
	&{[12 10 11 7 9 6 5 1 4 3 8 2]}
	&{[11 10 2 7 9 6 5 1 4 3 8]}
	&{[10 8 2 7 9 6 5 1 4 3]}
	&{[8 9 2 7 3 6 5 1 4]}
	&{[9 4 2 7 3 6 5 1]}
	&{[4 7 2 1 3 6 5]}
	*/
}
