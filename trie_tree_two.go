package search_target_world

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
				flag = true
				// fmt.Println(string(runes), "  here true...")
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

func in(runes []rune, backPosition, position int, allowBack bool, backParent, parent *NodeModelTwo) (ok bool, index int) {
	if parent == nil {
		if backPosition < len(runes) && allowBack && backParent != nil {
			// 回退
			// fmt.Println(string(runes[backPosition:]), "回退1 ", string(backParent.Character))
			return in(runes, backPosition+1, backPosition+1, allowBack, backParent, backParent)
		}
		return false, -1
	}
	if parent.IsPathEnd() {
		return true, position
	}
	// fmt.Println(allowBack, " ", string(runes[:position]), " ", string(parent.Character))
	if len(runes) <= 0 || position >= len(runes) {
		if backPosition < len(runes) && allowBack && backParent != nil {
			// 回退
			// fmt.Println("回退2")
			return in(runes, backPosition+1, backPosition+1, allowBack, backParent, backParent)
		}
		return false, -1
	}

	var current *NodeModelTwo
	if currentTemp, ok := parent.Children[ChildrenKey{
		Character:  runes[position],
		isContinue: true,
	}]; ok {
		current = currentTemp
		if position+1 < len(runes) {
			// 说明已然是最后一个元素了
			// return true, position
			// }

			nextTemp := runes[position+1]
			if _, ok := currentTemp.Children[ChildrenKey{
				Character:  nextTemp,
				isContinue: true,
			}]; ok {
				return in(runes, backPosition, position+1, allowBack, backParent, current)
			}
			if _, ok := current.Children[ChildrenKey{
				Character:  nextTemp,
				isContinue: false,
			}]; ok {
				return in(runes, backPosition, position+1, allowBack, backParent, current)
			}
		} else {
			if current.IsPathEnd() {
				return true, position + 1
			}
		}

	}
	if currentTemp, ok := parent.Children[ChildrenKey{
		Character:  runes[position],
		isContinue: false,
	}]; ok {
		current = currentTemp
		if position+1 < len(runes) {
			// 说明已然是最后一个元素了
			// return true, position
			// }
			// return in(runes, backPosition, position+1, false, parent, current)
			// fmt.Println("==")
			return in(runes, backPosition, position+1, false, parent, current)

			// return in(runes, backPosition, position+1, false, backParent, current)

			// nextTemp := runes[position+1]
			// if _, ok := current.Children[ChildrenKey{
			// 	Character:  nextTemp,
			// 	isContinue: true,
			// }]; ok {
			// return in(runes, backPosition, position+1, allowBack, backParent, current)
			// }

			// if _, ok := current.Children[ChildrenKey{
			// 	Character:  nextTemp,
			// 	isContinue: false,
			// }]; ok {
			// return in(runes, backPosition, position+1, allowBack, backParent, current)
			// }
		} else {
			if current.IsPathEnd() {
				// fmt.Println("is here ??")
				return true, position + 1
			}
		}
		// fmt.Println("可能会启用回退机制....")
	}
	if current == nil {
		current = parent
	}
	if current.isContinue {
		if backPosition < len(runes) && allowBack && backParent != nil {
			// 回退
			// fmt.Println(string(runes[backPosition:]), "回退3 ", "允许回退？：", allowBack, string(backParent.Character))

			return in(runes, backPosition+1, backPosition+1, allowBack, backParent, backParent)
		}
		return false, -1
	}
	return in(runes, backPosition, position+1, allowBack, backParent, current)
}

// TODO
// 缺少部分回退机制

// FindIn 检测关键字 -> 不连续 应用于规则1 和 2
func (tree *TrieModelTwo) FindIn(text string) (bool, string) {
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
				// fmt.Println("已然查找失败,寻找下一个可能存在的关键字")
				nowFound = false
				position = nowFoundPosition
				continue
			} else {
				break
			}
		}

		// 先看看有没有递归的必要
		isRunesContinue := false
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
						isRunesContinue = true
					}

					if _, ok := current.Children[ChildrenKey{
						Character:  nextTemp,
						isContinue: false,
					}]; ok {
						isRunesContinue = true
					}
				}

				// if current, ok := parent.Children[ChildrenKey{
				// 	Character:  currentTemp,
				// 	isContinue: false,
				// }]; ok {
				// 	nextTemp := runes[position+1]
				// 	if _, ok := current.Children[ChildrenKey{
				// 		Character:  nextTemp,
				// 		isContinue: true,
				// 	}]; ok {
				// 		isRunesContinue = true
				// 	}
				// 	if _, ok := current.Children[ChildrenKey{
				// 		Character:  nextTemp,
				// 		isContinue: false,
				// 	}]; ok {
				// 		isRunesContinue = true
				// 	}
				// }

			}
		}

		// if isRunesContinue == true 优先考虑
		// fmt.Println("isRunesContinue: ", isRunesContinue)
		current, found = parent.Children[ChildrenKey{
			Character:  runes[position],
			isContinue: isRunesContinue,
		}]
		// fmt.Println("found......: ", found)
		if found {
			allowBack := false
			if found && parent.isContinue == false {
				allowBack = true
				// fmt.Println("allowBack: ", allowBack, " parent.character:", string(parent.Character), " current.Character:", string(current.Character))
			}
			// TODO 递归到底 | 目前必须如此
			// fmt.Println("递归到底 | 目前必须如此....", allowBack)
			// fmt.Println(string(runes[:position]), " ", string(current.Character))
			ok, index := in(runes, position+1, position+1, allowBack, current, current)
			if ok {
				// fmt.Println("递归结果test......")
				return true, string(runes[left:index])
			} else {
				// fmt.Println("递归查找失败")
				if isRunesContinue {
					// 也可以尝试false
					// fmt.Println("也可以尝试false")
					if currentTemp, foundTemp := parent.Children[ChildrenKey{
						Character:  runes[position],
						isContinue: false,
					}]; foundTemp {
						// isContinue == false  --> allowBack=true
						allowBack := true
						// TODO 递归到底 | 目前必须如此
						ok, index := in(runes, position+1, position+1, allowBack, currentTemp, currentTemp)
						if ok {
							// fmt.Println("递归结果test......")
							return true, string(runes[left:index])
						} else {
							// fmt.Println("递归查找失败")
						}
					}

				}
			}
		}

		if !found || (!current.IsPathEnd() && position == length-1) {
			// fmt.Println("nowFound: ", nowFound)
			if !nowFound {
				parent = tree.Root
				position = left
				left++
				continue
			}

		}

		if found {
			if nowFound == false {
				nowFoundPosition = position
			}
			nowFound = true
			parent = current
		} else {
			if parent.isContinue {
				// fmt.Println("position: ", position, " left:", left)
				parent = tree.Root
				position = left
				left++
				continue
			}
		}
		if left <= position {
			// fmt.Println(found, " ? ", string(runes[left:position+1]))
		}
		if found && current.IsPathEnd() && left <= position {
			// TODO 目前返回的string可能不正确，需要重新调整！
			return true, string(runes[left : position+1])
		}
	}
	// fmt.Println("return last.....")

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
