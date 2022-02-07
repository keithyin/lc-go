package algos

type TrieNode struct {
	canBeEnd bool
	edges    []*TrieEdge
}

func NewTrieNode() *TrieNode {
	trieNode := new(TrieNode)
	trieNode.edges = make([]*TrieEdge, 0)
	return trieNode
}

type TrieEdge struct {
	charcater byte
	node      *TrieNode
}

func newTrieEdge(char byte) *TrieEdge {
	trieEdge := new(TrieEdge)
	trieEdge.charcater = char
	trieEdge.node = NewTrieNode()
	return trieEdge
}

type Trie struct {
	root *TrieNode
}

func TrieConstructor() Trie {
	return Trie{root: NewTrieNode()}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, char := range []byte(word) {
		found := false
		var nextNode *TrieNode
		for _, edge := range cur.edges {
			if edge.charcater == char {
				found = true
				nextNode = edge.node
			}
		}
		if found {
			cur = nextNode
		} else {
			newEdge := newTrieEdge(char)
			cur.edges = append(cur.edges, newEdge)
			cur = newEdge.node
		}
	}
	cur.canBeEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, char := range []byte(word) {
		found := false
		var nextNode *TrieNode
		for _, edge := range cur.edges {
			if edge.charcater == char {
				found = true
				nextNode = edge.node
			}
		}
		if found {
			cur = nextNode
		} else {
			return false
		}
	}
	return cur.canBeEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, char := range []byte(prefix) {
		found := false
		var nextNode *TrieNode
		for _, edge := range cur.edges {
			if edge.charcater == char {
				found = true
				nextNode = edge.node
			}
		}
		if found {
			cur = nextNode
		} else {
			return false
		}
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
