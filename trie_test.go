package trie

import (
	"reflect"
	"testing"
)

func prepareTrie() *Trie {
	t := NewTrie()
	for _, word := range []string{"hello", "hell", "hamster", "harris", "harrier", "harmore", "harmless"} {
		t.Insert(word)
	}
	return t
}

func TestTrie_HasPrefix(t *testing.T) {
	tree := prepareTrie()

	if tree.HasPrefix("he") != true {
		t.Fatal("tree.HasPrefix(\"he\") should be true")
	}
	if tree.HasPrefix("") != true {
		t.Fatal("tree.HasPrefix(\"\") should be true")
	}

	if tree.HasPrefix("hem") != false {
		t.Fatal("tree.HasPrefix(\"\") should be false")
	}
}

func TestTrie_HasWord(t *testing.T) {
	tree := prepareTrie()

	if tree.HasWord("hello") != true {
		t.Fatal("tree.HasWord(\"hello\") should be true")
	}
	if tree.HasWord("hamster") != true {
		t.Fatal("tree.HasWord(\"hamster\") should be true")
	}

	if tree.HasWord("hem") != false {
		t.Fatal("tree.HasWord(\"hem\") should be false")
	}
	if tree.HasWord("hel") != false {
		t.Fatal("tree.HasWord(\"hel\") should be false")
	}
	if tree.HasWord("") != false {
		t.Fatal("tree.HasWord(\"\") should be false")
	}
}

func TestTrie_Insert(t *testing.T) {
	tree := prepareTrie()
	tree.Insert("inserted")

	if tree.HasWord("inserted") != true {
		t.Fatal("tree.HasWord(\"inserted\") should be true")
	}
}

func TestTrie_WordsByPrefix(t *testing.T) {
	tree := prepareTrie()
	words := tree.WordsByPrefix("he")
	expected := []string{"hell", "hello"}

	if !reflect.DeepEqual(words, expected) {
		t.Fatalf("words should be equal to %v", expected)
	}

	words = tree.WordsByPrefix("gag")
	if len(words) > 0 {
		t.Fatal("There are no words with prefix gag")
	}
}

func BenchmarkTrie_HasWord(b *testing.B) {
	t := prepareTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t.HasWord("harris")
	}
}

func BenchmarkTrie_HasPrefix(b *testing.B) {
	t := prepareTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t.HasPrefix("harm")
	}
}

func BenchmarkTrie_WordsByPrefix(b *testing.B) {
	t := prepareTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t.WordsByPrefix("ha")
	}
}
