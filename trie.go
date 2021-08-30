package trie

// Trie represents trie tree data structure
//
// Use NewTrie to create new tree
type Trie struct {
	root *trieNode
}

type trieNode struct {
	isWord   bool
	children map[rune]*trieNode
}

// NewTrie creates new tree along with its root node
func NewTrie() *Trie {
	return &Trie{root: &trieNode{children: map[rune]*trieNode{}}}
}

// Insert inserts word to the tree. Each letter becomes tree node (if not exists)
func (t *Trie) Insert(word string) {
	current := t.root
	for _, letter := range word {
		node, ok := current.children[letter]
		if !ok {
			node = &trieNode{children: map[rune]*trieNode{}}
			current.children[letter] = node
		}
		current = node
	}
	current.isWord = true
}

func (t *trieNode) nodeByPrefix(word string) *trieNode {
	var ok bool
	current := t
	for _, letter := range word {
		current, ok = current.children[letter]
		if !ok {
			return nil
		}
	}
	return current
}

// HasWord checks if the tree has word. It needs the complete match. For prefix check use HasPrefix
func (t *Trie) HasWord(word string) bool {
	node := t.root.nodeByPrefix(word)
	if node == nil {
		return false
	}
	return node.isWord
}

// HasPrefix check if the tree has words with this prefix
func (t *Trie) HasPrefix(word string) bool {
	node := t.root.nodeByPrefix(word)
	if node == nil {
		return false
	}
	return true
}

// WordsByPrefix returns all words from the tree with this prefix
func (t *Trie) WordsByPrefix(prefix string) []string {
	return t.root.wordsByPrefix(prefix)
}

func (t *trieNode) wordsByPrefix(prefix string) []string {
	node := t.nodeByPrefix(prefix)
	if node == nil {
		return nil
	}
	return node.travers(prefix)
}

type traversHelper struct {
	node   *trieNode
	prefix string
}

type stack []*traversHelper

func (s *stack) push(th *traversHelper) {
	*s = append(*s, th)
}

func (s *stack) pop() *traversHelper {
	if len(*s) > 0 {
		n := len(*s) - 1
		elem := (*s)[n]
		*s = (*s)[:n]
		return elem
	}
	return nil
}

func (t *trieNode) travers(prefix string) []string {
	var words []string
	stack := stack{{t, prefix}}
	for {
		th := stack.pop()
		if th == nil {
			return words
		}
		for letter, child := range th.node.children {
			if child.isWord {
				words = append(words, th.prefix+string(letter))
			}
			if len(child.children) > 0 {
				stack.push(&traversHelper{child, th.prefix + string(letter)})
			}
		}
	}
}
