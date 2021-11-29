package tests

import (
	"testing"

	trie "github.com/bakyazi/goalgo/trie-str-search"
)

func TestTrie(t *testing.T) {
	head := &trie.TrieNode{}
	texts := []string{"berkay", "brky", "berky"}

	for _, t := range texts {
		head.Insert(t)
	}

	if !head.LookUp("aberkay") {
		t.Log("aberkay")
		t.Fail()
	}
	if !head.LookUp("dasdabrky") {
		t.Log("dasdabrky")
		t.Fail()
	}
	if !head.LookUp("berkyfsdfsdfsd") {
		t.Log("berkyfsdfsdfsd")
		t.Fail()
	}
	if head.LookUp("bdedrkayyu") {
		t.Log("bdedrkayyu")
		t.Fail()
	}

	if head.LookUp("ber") {
		t.Log("ber")
		t.Fail()
	}
}
