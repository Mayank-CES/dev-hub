package main

import "fmt"

// ArraySize is the size of the hash table array
const ARRAY_SIZE = 7

// HashTable will hold an array
type HashTable struct {
	array [ARRAY_SIZE]*bucket
}

// bucker is a linked list in each slot of the hash table array
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert will take in a create a node with the same key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		node := &bucketNode{key: k}
		node.next = b.head
		b.head = node
	} else {
		fmt.Println("already exists")
	}
}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take a key and delete the node
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ARRAY_SIZE
}

// New will create a bucket in each slot of the hash table
func New() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	bucket := &bucket{}
	bucket.insert("MAYANK")
	bucket.insert("MAYANK")              // Output: already exists
	fmt.Println(bucket.search("MAYANK")) // Output: true
	fmt.Println(bucket.search("SAKETH")) // Output: false

	bucket.delete("MAYANK")
	fmt.Println(bucket.search("MAYANK")) // Output: false

	hashTable := &HashTable{}
	fmt.Println(hashTable) // Output: &{[<nil> <nil> <nil> <nil> <nil> <nil> <nil>]}
	hashTable = New()
	fmt.Println(hashTable)      // Output: &{[0xc00004a020 0xc00004a028 0xc00004a030 0xc00004a038 0xc00004a040 0xc00004a048 0xc00004a050]}
	fmt.Println(hash("MAYANK")) // Output: 1
	list := []string{
		"MAYANK",
		"UNNAT",
		"PULKIT",
		"UTKARSH",
		"SAKETH",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("MAYANK")
	fmt.Println(hashTable.Search("MAYANK")) // Output: false
	fmt.Println(hashTable.Search("SAKETH")) // Output: true
}
