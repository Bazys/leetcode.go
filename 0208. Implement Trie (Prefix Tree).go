package leetcode

type Trie struct {
	next   map[int32]*Trie
	isLeaf bool
}

func NewTrie() Trie {
	return Trie{
		next:   make(map[int32]*Trie),
		isLeaf: false}
}

func (t *Trie) Insert(word string) {
	cur := t
	for _, ch := range word {
		if _, ok := cur.next[ch]; !ok {
			cur.next[ch] = &Trie{
				next:   make(map[int32]*Trie),
				isLeaf: false,
			}
		}
		cur = cur.next[ch]
	}
	cur.isLeaf = true
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, ch := range word {
		if _, ok := cur.next[ch]; !ok {
			return false
		}
		cur = cur.next[ch]
	}
	return cur.isLeaf
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for _, ch := range prefix {
		if _, ok := cur.next[ch]; !ok {
			return false
		}
		cur = cur.next[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
