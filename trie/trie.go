package trie

const (
	alphabetSize = 26
)

type trieNode struct {
	children  [alphabetSize]*trieNode
	isWordEnd bool
}

// Trie, tree data structure: https://en.wikipedia.org/wiki/Trie.
// Based on https://golangbyexample.com/trie-implementation-in-go/.
type trie struct {
	root *trieNode
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'

		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}

		current = current.children[index]
	}

	current.isWordEnd = true
}

func (t *trie) find(word string) bool {
	wordLength := len(word)
	current := t.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'

		if current.children[index] == nil {
			return false
		}

		current = current.children[index]
	}

	return current.isWordEnd
}

func main() {

}
