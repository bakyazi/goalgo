# trie-str-search

Check a string contains a bucket of string or not.

It may be used for whether a string contains any muted word or not


Insert a muted word into Trie
```go

import (
    trie "github.com/bakyazi/goalgo/trie-str-search"
)

func main() {
    bannedWords := &trie.TrieNode{}

    bannedWords.Insert("muted")
    bannedWords.Insert("mted")
    bannedWords.Insert("mued")
    bannedWords.Insert("uted")
    
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
  /      /           \     \
 Ø      Ø              d     Ø
                        \
                         Ø
```

** Ø is end of muted word

Check whether a string contains any muted words or not:

```go
bannedWords.LookUp("fsdmuted432423") // true
bannedWords.LookUp("something") // false
bannedWords.LookUp("mute") // false
```
