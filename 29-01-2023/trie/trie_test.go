package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	// Test insert
	trie.Insert("hello")
	trie.Insert("hell")
	trie.Insert("he")

	// Test search
	tests := []struct {
		word     string
		expected bool
	}{
		{"hello", true},
		{"hell", true},
		{"he", true},
		{"hel", false},
		{"", false},
	}
	for _, test := range tests {
		if trie.Search(test.word) != test.expected {
			t.Errorf("Search(%q) = %v, expected %v", test.word, trie.Search(test.word), test.expected)
		}
	}

	// Test startsWith
	prefix_tests := []struct {
		prefix   string
		expected bool
	}{
		{"hello", true},
		{"hell", true},
		{"he", true},
		{"het", false},
		{"", true},
	}
	for _, test := range prefix_tests {
		if trie.StartsWith(test.prefix) != test.expected {
			t.Errorf("StartsWith(%q) = %v, expected %v", test.prefix, trie.StartsWith(test.prefix), test.expected)
		}
	}
}
