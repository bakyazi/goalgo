package trie

type Children map[rune]*TrieNode
type TrieNode struct {
	Children Children
}

func (t *TrieNode) Insert(text string) {
	t.insertChars([]rune(text))
}

func (t *TrieNode) insertChars(chars []rune) {
	if t.Children == nil {
		t.Children = make(Children)
	}
	if len(chars) == 0 {
		t.Children[0] = &TrieNode{}
		return
	}
	char := chars[0]
	if node, ok := t.Children[char]; ok {
		node.insertChars(chars[1:])
	} else {
		newNode := TrieNode{}
		t.Children[char] = &newNode
		t.Children[char].insertChars(chars[1:])
	}
}
func (t *TrieNode) IsContains(chars []rune) bool {
	if _, ok := t.Children[0]; ok {
		// reach end of one of searching words
		return true
	} else if len(chars) == 0 {
		// not reach end of any word but any letter isn't left
		return false
	}
	if node, ok := t.Children[chars[0]]; !ok {
		return false
	} else {
		return node.IsContains(chars[1:])
	}
}

func (t TrieNode) LookUp(s string) bool {
	chars := []rune(s)

	for len(chars) > 0 {
		if t.IsContains(chars) {
			return true
		}
		chars = chars[1:]
	}
	return false
}
