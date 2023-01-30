package trie

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func Constructor() Trie {
	return Trie{&TrieNode{children: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		if node.children == nil {
			node.children = make(map[rune]*TrieNode)
		}
		child, ok := node.children[char]
		if !ok {
			child = &TrieNode{children: make(map[rune]*TrieNode)}
			node.children[char] = child
		}
		node = child
	}
	node.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, char := range word {
		child, ok := node.children[char]
		if !ok {
			return false
		}
		node = child
	}
	return node.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, char := range prefix {
		child, ok := node.children[char]
		if !ok {
			return false
		}
		node = child
	}
	return true
}
