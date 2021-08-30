## go-trie

Package go-trie implements trie tree (or prefix tree) data structure useful
for things like prefix search/autocompletion.

For now, it supports Insert, HasWord, HasPrefix and WordsByPrefix methods. 
WordsByPrefix collects all words with given prefix without usage of recursion while search.

```go
t := NewTrie()

t.Insert("go", "golang", "gopher", "python", "pythonista", "grow", "gg", "glitch", "glass")

t.HasPrefix("gol") // return true
t.HasPrefix("gole") // return false

t.HasWord("gopher") // return true
t.HasWord("oh") // return false

words := t.WordsByPrefix("go") // return []string{"go", "gopher", "golang"}
```
