package top100

import "testing"

func Test_lc208(t *testing.T) {
	//
	trie := TrieConstructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("expect true get false")
	}
	if trie.Search("app") {
		t.Errorf("expect false get true")
	}
	if !trie.StartsWith("app") {
		t.Errorf("expect true get false")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("expect true get false")
	}
}
