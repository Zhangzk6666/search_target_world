package search_target_world

// Trie 短语组成的Trie树.
type Trie struct {
	Root *Node
}

// Node Trie树上的一个节点.
type Node struct {
	isRootNode bool
	isPathEnd  bool
	Character  rune
	Children   map[rune]*Node
}

// NewTrie 新建一棵Trie
func NewTrie() *Trie {
	return &Trie{
		Root: NewRootNode(0),
	}
}

// Add 添加若干个词
func (tree *Trie) Add(words ...string) {
	for _, word := range words {
		tree.add(word)
	}
}

func (tree *Trie) add(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.Children[r]; ok {
			current = next
		} else {
			newNode := NewNode(r)
			current.Children[r] = newNode
			current = newNode
		}
		if position == len(runes)-1 {
			current.isPathEnd = true
		}
	}
}

func (tree *Trie) Del(words ...string) {
	for _, word := range words {
		tree.del(word)
	}
}

func (tree *Trie) del(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.Children[r]; !ok {
			return
		} else {
			current = next
		}

		if position == len(runes)-1 {
			current.SoftDel()
		}
	}
}

// FindIn 检测关键字
func (tree *Trie) FindIn(text string) (bool, string) {
	const (
		Empty = ""
	)
	var (
		parent  = tree.Root
		current *Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < len(runes); position++ {
		current, found = parent.Children[runes[position]]

		if !found || (!current.IsPathEnd() && position == length-1) {
			parent = tree.Root
			position = left
			left++
			continue
		}

		if current.IsPathEnd() && left <= position {
			return true, string(runes[left : position+1])
		}

		parent = current
	}

	return false, Empty
}

// NewNode 新建子节点
func NewNode(character rune) *Node {
	return &Node{
		Character: character,
		Children:  make(map[rune]*Node, 0),
	}
}

// NewRootNode 新建根节点
func NewRootNode(character rune) *Node {
	return &Node{
		isRootNode: true,
		Character:  character,
		Children:   make(map[rune]*Node, 0),
	}
}

// IsPathEnd 判断是否为某个路径的结束
func (node *Node) IsPathEnd() bool {
	return node.isPathEnd
}

// SoftDel 置软删除状态
func (node *Node) SoftDel() {
	node.isPathEnd = false
}
