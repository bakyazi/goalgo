package trie

type Children map[rune]*TrieNode
type TrieNode struct {
	Children Children
}

func (t *TrieNode) Insert(text string) {
	t.insertChars([]rune(text))
}

func (t *TrieNode) insertChars(chars []rune) {
	if len(chars) == 0 {
		return
	}

	if t.Children == nil {
		t.Children = make(Children)
	}

	char := chars[0]
	if node, ok := t.Children[char]; ok {
		node.insertChars(chars[1:])
	} else {
		newNode := TrieNode{Children: make(Children)}
		t.Children[char] = &newNode
		t.Children[char].insertChars(chars[1:])
	}
}
func (t *TrieNode) IsContains(chars []rune) bool {
	if len(t.Children) == 0 {
		return true
	}
	if node, ok := t.Children[chars[0]]; !ok {
		return false
	} else {
		if len(chars[1:]) == 0 {
			return true
		}
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
