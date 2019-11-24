package treeformation

func CreateNode(nodeData int64) *Node {
	return &Node{
		Data:  nodeData,
		Left:  nil,
		Right: nil,
	}
}

func CreateTree(datalist []int64) *Node {
	stack := []*Node{}
	var head, temphead *Node
	var i int
	dataLength := len(datalist)
	for i < len(datalist) {
		if head == nil {
			head = CreateNode(datalist[i])
			stack = append(stack, head)
			i++
		} else {
			temphead = stack[0]
			if i < dataLength {
				if datalist[i] == 0 {
					if temphead != nil {
						temphead.Left = nil
					}
					i++
				} else {
					temphead.Left = CreateNode(datalist[i])
					stack = append(stack, temphead.Left)
					i++
				}
			}
			if i < dataLength {
				if datalist[i] == 0 {
					if temphead != nil {
						temphead.Right = nil
					}
					i++
				} else {
					temphead.Right = CreateNode(datalist[i])
					stack = append(stack, temphead.Right)
					i++
				}
			}
			stack = stack[1:]
		}
	}
	return head
}


func CreateNodeNextRight(nodeData int64) *NodeWihNextRight {
	return &NodeWihNextRight{
		Data:  nodeData,
		Left:  nil,
		Right: nil,
		NextRight:nil,
	}
}

func CreateTreeNextRight(datalist []int64) *NodeWihNextRight {
	stack := []*NodeWihNextRight{}
	var head, temphead *NodeWihNextRight
	var i int
	dataLength := len(datalist)
	for i < len(datalist) {
		if head == nil {
			head = CreateNodeNextRight(datalist[i])
			stack = append(stack, head)
			i++
		} else {
			temphead = stack[0]
			if i < dataLength {
				if datalist[i] == 0 {
					if temphead != nil {
						temphead.Left = nil
					}
					i++
				} else {
					temphead.Left = CreateNodeNextRight(datalist[i])
					stack = append(stack, temphead.Left)
					i++
				}
			}
			if i < dataLength {
				if datalist[i] == 0 {
					if temphead != nil {
						temphead.Right = nil
					}
					i++
				} else {
					temphead.Right = CreateNodeNextRight(datalist[i])
					stack = append(stack, temphead.Right)
					i++
				}
			}
			stack = stack[1:]
		}
	}
	return head
}

