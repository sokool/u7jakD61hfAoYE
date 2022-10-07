package dictionary

import (
	"sync"
	"unicode"
)

// Storage is prefix tree - it represents letters on every node, where deepest node in path is marked as word
type Storage struct {
	mu sync.RWMutex
	*node
}

func NewStorage(w ...Word) *Storage {
	return (&Storage{node: newNode()}).Put(w...)
}

func (s *Storage) Put(w ...Word) *Storage {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range w {
		c := s.node
		for k := range w[i] {
			l := unicode.ToLower(rune(w[i][k]))
			if _, ok := c.letters[l]; !ok {
				c.letters[l] = newNode()
			}
			c = c.letters[l]
		}
		c.count++
	}
	return s
}

func (s *Storage) Exists(w Word) bool {
	n := s.read(w)
	if n == nil {
		return false
	}
	return n.count > 0
}

func (s *Storage) Search(w Word) Words {
	a, n := make(Words), s.read(w)
	if n == nil {
		return a
	}
	if n.count > 0 {
		a[w] = n.count
		return a
	}
	n.find(w, a)
	return a
}

func (s *Storage) read(w Word) *node {
	s.mu.RLock()
	defer s.mu.RUnlock()

	n := s.node
	for i := range w {
		l := unicode.ToLower(rune(w[i]))
		if n.letters[l] == nil {
			return nil
		}
		n = n.letters[l]
	}

	return n
}

// node contains map of rune instead array of 26 [byte]*node due reason that it's easy to adapt this prefix tree to
// support unicode characters such as ∂łóćź... instead only english letters.
// Sure, it cost more memory allocations when nodes are added to the prefix tree but it gives possibility to store any
// word from any language.
type node struct {
	letters map[rune]*node
	// count when non-zero node is marked as end of the word, when more than 1 then word was stored multiple times
	count int
}

func newNode() *node {
	return &node{map[rune]*node{}, 0}
}

// find uses DFS in order to find Words from selected node
func (n *node) find(prefix Word, w Words) {
	if n.count > 0 {
		w[prefix] = n.count
	}
	for l, c := range n.letters {
		c.find(prefix.append(l), w)
	}
}
