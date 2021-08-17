package main

import "fmt"

const MAXCAP = 26

type Trie struct {
	isWord bool
	next   map[rune]*Trie
}

// Initialize your data structure here.
func Constructor() Trie {
	return Trie{
		isWord: false,
		next:   make(map[rune]*Trie, MAXCAP),
	}
}

// Inserts a word into the trie.
func (t *Trie) insert(word string) {
	for _, v := range word {
		if t.next[v] == nil {
			node := Trie{
				isWord: false,
				next:   make(map[rune]*Trie, MAXCAP),
			}
			t.next[v] = &node
		}
		t = t.next[v]
	}
	t.isWord = true
}

// Returns if the word is in the trie.
func (t *Trie) search(word string) bool {
	for _, v := range word {
		if t.next[v] == nil {
			return false
		}
		t = t.next[v]
	}
	return t.isWord
}

// Returns if there is any word in the trie that starts woth the given prefix
func (t *Trie) startsWith(prefix string) bool {
	for _, v := range prefix {
		if t.next[v] == nil {
			return false
		}
		t = t.next[v]
	}
	return t.isWord
}

func main() {
	t := Constructor()
	t.insert("Hello")
	fmt.Println(t.search("Hello"))
	fmt.Println(t.search("Hallo"))

}
