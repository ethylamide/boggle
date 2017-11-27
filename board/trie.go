package board

type Trie struct {
	vertex   map[rune]*Trie
	symbol   rune
	end      bool
	repeated bool
}

func (t *Trie) AddString(s []rune) {
	trie := t
	for _, r := range s {
		if trie.vertex == nil {
			trie.vertex = make(map[rune]*Trie)
		}

		if subTrie, ok := trie.vertex[r]; ok {
			trie = subTrie
		} else {
			trie.vertex[r] = &Trie{symbol: r, end: false, repeated: false}
			trie = trie.vertex[r]
		}
	}
	trie.end = true
	trie.repeated = false
}

func (t *Trie) Contains(s []rune) bool {
	trie := t
	for _, r := range s {
		if trie.vertex == nil {
			return false
		}

		if subTrie, ok := trie.vertex[r]; ok {
			trie = subTrie
		} else {
			return false
		}
	}

	if trie.repeated || !trie.end {
		return false
	}
	trie.repeated = true
	return true
}

func (t *Trie) StartWith(s []rune) bool {
	trie := t
	for _, r := range s {
		if trie.vertex == nil {
			return false
		}

		if subTrie, ok := trie.vertex[r]; ok {
			trie = subTrie
		} else {
			return false
		}
	}
	return true
}
