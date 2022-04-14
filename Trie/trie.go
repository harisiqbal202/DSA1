package main

import "fmt"

const nodeChildrenSize = 26

type Node struct {
	children [nodeChildrenSize]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}
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
func (t *Trie) Serach(w string) bool {
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
	fmt.Println("Tries Implements")
	testTrie := InitTrie()
	fmt.Println(testTrie.root)
	testTrie.Insert("aragorn")
	fmt.Println(testTrie.Serach("aragorn"))
}
