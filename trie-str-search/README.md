# trie-str-search

Check a string contains a bucket of string or not.

It may be used for whether a string contains any muted word or not


Insert a muted word into Trie
```go

import (
    trie "github.com/bakyazi/goalgo/trie-str-search"
)

func main() {
    trieNode := &trie.TrieNode{}

    trieNode.Insert("muted")
    trieNode.Insert("mted")
    trieNode.Insert("mued")
    
}
```

It generate a trie like below

```
            head
            /  \
           m    u
         /   \    \
        t     u     t
       /     / \     \
      e     e   t     e 
    /      /      \     \
   d      d        e     d
                    \
                     d
```

Check a string contains any muted words

```go
trieNode.LookUp("fsdmuted432423")
```
