## go-trie

Package go-trie implements trie tree (or prefix tree) data structure useful
for things like prefix search/autocompletion.

For now, it supports Insert, HasWord, HasPrefix and WordsByPrefix methods. 
WordsByPrefix collects all words with given prefix without usage of recursion while search.

```go
tree := NewTrie()

t.Insert("go", "golang", "gopher", "python", "pythonista", "grow", "gg", "glitch", "glass")

hasPrefix := tree.HasPrefix("gol") // return true
hasPrefix = tree.HasPrefix("gene") // return false

hasWord := tree.HasWord("gopher") // return true
hasWord = tree.HasWord("foo") // return false

words := tree.WordsByPrefix("go") // return []string{"go", "gopher", "golang"}
```
