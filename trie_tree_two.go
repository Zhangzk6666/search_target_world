package search_target_world

import "fmt"

// TrieModelTwo 短语组成的TrieModelTwo树.
type TrieModelTwo struct {
	Root *NodeModelTwo
}

// NodeModelTwo TrieModelTwo树上的一个节点.
type NodeModelTwo struct {
	isRootNodeModelTwo bool
	isPathEnd          bool
	isContinue         bool
	Character          rune
	Children           map[ChildrenKey]*NodeModelTwo
	// Children           map[rune]*NodeModelTwo

}
type ChildrenKey struct {
	Character  rune
	isContinue bool
}

// NewTrieModelTwo 新建一棵TrieModelTwo
func NewTrieModelTwo() *TrieModelTwo {
	return &TrieModelTwo{
		Root: NewRootNodeModelTwo(0),
	}
}

// Add 添加若干个词
func (tree *TrieModelTwo) Add(words ...string) {
	for _, word := range words {
		tree.add(word)
	}
}

func (tree *TrieModelTwo) add(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if r == []rune("|")[0] {
			continue
		}
		flag := false
		if position+1 < len(runes) {
			if runes[position+1] != []rune("|")[0] {
				fmt.Println("here true.....")
				fmt.Println(word, " ", position)
				flag = true
			}
		}
		if next, ok := current.Children[ChildrenKey{
			Character:  r,
			isContinue: flag,
		}]; ok && next.isContinue == flag {
			current = next
		} else {
			newNodeModelTwo := NewNodeModelTwo(r)
			newNodeModelTwo.isContinue = flag
			current.Children[ChildrenKey{
				Character:  r,
				isContinue: flag,
			}] = newNodeModelTwo
			current = newNodeModelTwo
		}
		if position == len(runes)-1 {
			current.isPathEnd = true
		}
	}
}

// func (tree *TrieModelTwo) Del(words ...string) {
// 	for _, word := range words {
// 		tree.del(word)
// 	}
// }

// TODO isContinue !!
// func (tree *TrieModelTwo) del(word string) {
// 	var current = tree.Root
// 	var runes = []rune(word)
// 	for position := 0; position < len(runes); position++ {
// 		r := runes[position]
// 		if next, ok := current.Children[ChildrenKey{
// 			Character:  r,
// 			isContinue: flag,
// 		}]; !ok {
// 			return
// 		} else {
// 			current = next
// 		}

// 		if position == len(runes)-1 {
// 			current.SoftDel()
// 		}
// 	}
// }

// TODO
func in(runes []rune, position int, parent *NodeModelTwo) (ok bool, index int) {
	fmt.Println("position:  ", position)
	if parent == nil {
		return false, -1
	}
	if parent.IsPathEnd() {
		return true, position
	}
	if len(runes) <= 0 || position >= len(runes) {
		return false, -1
	}

	var current *NodeModelTwo
	if currentTemp, ok := parent.Children[ChildrenKey{
		Character:  runes[position],
		isContinue: true,
	}]; ok {
		current = currentTemp
		nextTemp := runes[position+1]
		if next, ok := currentTemp.Children[ChildrenKey{
			Character:  nextTemp,
			isContinue: true,
		}]; ok {
			return in(runes, position+1, next)
		}
		if next, ok := current.Children[ChildrenKey{
			Character:  nextTemp,
			isContinue: false,
		}]; ok {
			return in(runes, position+1, next)
		}
	}
	if currentTemp, ok := parent.Children[ChildrenKey{
		Character:  runes[position],
		isContinue: false,
	}]; ok {
		current = currentTemp
		nextTemp := runes[position+1]
		if next, ok := current.Children[ChildrenKey{
			Character:  nextTemp,
			isContinue: true,
		}]; ok {
			return in(runes, position+1, next)
		}

		if next, ok := current.Children[ChildrenKey{
			Character:  nextTemp,
			isContinue: false,
		}]; ok {
			return in(runes, position+1, next)
		}
	}
	if current == nil {
		current = parent
	}
	if current.isContinue {
		return false, -1
	}
	return in(runes, position+1, current)
}

// FindIn 检测关键字 -> 不连续
func (tree *TrieModelTwo) FindInWithoutStrict(text string) (bool, string) {
	fmt.Println("FindInWithoutStrict()")
	const (
		Empty = ""
	)
	var (
		parent           = tree.Root
		current          *NodeModelTwo
		runes            = []rune(text)
		length           = len(runes)
		left             = 0
		found            bool
		nowFound         bool
		nowFoundPosition int
	)

	for position := 0; position <= len(runes); position++ {
		if position == len(runes) {
			if nowFound {
				// 已然查找失败,寻找下一个可能存在的关键字
				nowFound = false
				position = nowFoundPosition
				continue
			} else {
				break
			}
		}

		// 先看看有没有递归的必要
		isMust := false
		if position+1 < len(runes) {
			currentTemp := runes[position]
			if parent != nil {
				if current, ok := parent.Children[ChildrenKey{
					Character:  currentTemp,
					isContinue: true,
				}]; ok {
					nextTemp := runes[position+1]
					if _, ok := current.Children[ChildrenKey{
						Character:  nextTemp,
						isContinue: true,
					}]; ok {
						isMust = true
					}

					if _, ok := current.Children[ChildrenKey{
						Character:  nextTemp,
						isContinue: false,
					}]; ok {
						isMust = true
					}
				}

				if current, ok := parent.Children[ChildrenKey{
					Character:  currentTemp,
					isContinue: false,
				}]; ok {
					nextTemp := runes[position+1]
					if _, ok := current.Children[ChildrenKey{
						Character:  nextTemp,
						isContinue: true,
					}]; ok {
						isMust = true
					}

					if _, ok := current.Children[ChildrenKey{
						Character:  nextTemp,
						isContinue: false,
					}]; ok {
						isMust = true
					}
				}
			}
		}
		current, found = parent.Children[ChildrenKey{
			Character:  runes[position],
			isContinue: isMust,
		}]
		if !found || (!current.IsPathEnd() && position == length-1) {
			if !nowFound {
				parent = tree.Root
				position = left
				left++
				continue
			}
		}

		// TODO 递归到底 | 目前必须如此
		ok, index := in(runes, position+1, current)
		if ok {
			fmt.Println("递归结果test......")
			return true, string(runes[left:index])
		}

		if found {
			if nowFound == false {
				nowFoundPosition = position
			}
			nowFound = true
			parent = current
		}
		if left <= position {
			fmt.Println(found, " ? ", string(runes[left:position+1]))
		}
		if found && current.IsPathEnd() && left <= position {
			// TODO 目前返回的string可能不正确，需要重新调整！
			return true, string(runes[left : position+1])
		}
	}
	return false, Empty
}

// NewNodeModelTwo 新建子节点
func NewNodeModelTwo(character rune) *NodeModelTwo {
	return &NodeModelTwo{
		Character: character,
		Children:  make(map[ChildrenKey]*NodeModelTwo, 0),
	}
}

// NewRootNodeModelTwo 新建根节点
func NewRootNodeModelTwo(character rune) *NodeModelTwo {
	return &NodeModelTwo{
		isRootNodeModelTwo: true,
		Character:          character,
		Children:           make(map[ChildrenKey]*NodeModelTwo, 0),
	}
}

// IsPathEnd 判断是否为某个路径的结束
func (NodeModelTwo *NodeModelTwo) IsPathEnd() bool {
	return NodeModelTwo.isPathEnd
}

// SoftDel 置软删除状态
func (NodeModelTwo *NodeModelTwo) SoftDel() {
	NodeModelTwo.isPathEnd = false
}
