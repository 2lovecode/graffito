package top100

type Trie struct {
	children [26]*Trie
	end      bool
}

func TrieConstructor() Trie {
	return Trie{
		children: [26]*Trie{},
		end:      false,
	}
}

func (this *Trie) Insert(word string) {
	current := this

	for _, alpha := range word {
		idx := alpha - 'a'
		if current.children[idx] == nil {
			current.children[idx] = &Trie{}
		}
		current = current.children[idx]
	}
	current.end = true
}

func (this *Trie) Search(word string) bool {
	node := this.SearchWithPrefix(word)

	if node != nil && node.end {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchWithPrefix(prefix) != nil
}

func (this *Trie) SearchWithPrefix(prefix string) *Trie {
	if len(prefix) == 0 {
		return nil
	}
	current := this
	for _, alpha := range prefix {
		idx := alpha - 'a'
		if current.children[idx] == nil {
			return nil
		} else {
			current = current.children[idx]
		}
	}
	return current
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
