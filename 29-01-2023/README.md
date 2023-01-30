# Trie Implementation and Unit Test in Go
This is a solution to the problem of implementing a Trie and writing unit tests for the same, in Go.

## Problem Statement
The problem statement can be found at [LeetCode](https://leetcode.com/problems/implement-trie-prefix-tree)

## Implementation
The Trie implementation includes the following:

1. A Trie struct that contains a pointer to the root node.
2. A TrieNode struct that represents a node in the Trie and contains:
3. A map[rune]*TrieNode to store its children nodes
4. A bool flag isEndOfWord to indicate if the node represents the end of a word
5. A Constructor function that returns an instance of the Trie struct with a root node.
6. An Insert function that inserts a word into the Trie.
7. A Search function that searches for a word in the Trie.
8. A StartsWith function that searches for a prefix in the Trie.

## Unit Tests
The implementation includes unit tests for the Trie struct, to verify its functionality. The tests cover the following scenarios:

Inserting words into the Trie
Searching for words and prefixes in the Trie

## Note
This is the first problem that I have solved in Go.

## Running the Code
To run the code, you will need to have Go installed on your machine. You can download and install it from the official website: https://golang.org/

Once you have Go installed, you can run the tests by navigating to the directory containing the code and running the following command in your terminal:

```go
go test -v
```


