package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a poinnter to the root mode
type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and Return true is that word is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := NewTrie()
	toAdd := []string{
		"mayank",
		"data",
		"structure",
		"lemon",
		"lemonade",
	}
	for _, str := range toAdd {
		myTrie.Insert(str)
	}
	fmt.Println("Search: ", myTrie.Search("mayank"))
}
