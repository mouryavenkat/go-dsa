package main

import (
	"go-dsa/treeformation"
)

func insertNodeAtFirstVacantPos(headNode *treeformation.Node, nodeToInsert int) *treeformation.Node {
	stack := []*treeformation.Node{headNode}
	for len(stack) != 0 {
		element := stack[0]
		if element.Left == nil {
			newNode := treeformation.CreateNode(int64(nodeToInsert))
			element.Left = newNode
			break
		}
		if element.Right == nil {
			newNode := treeformation.CreateNode(int64(nodeToInsert))
			element.Right = newNode
			break
		}
		stack = append(stack, []*treeformation.Node{element.Left, element.Right}...)
		stack = stack[1:]
	}
	return headNode
}

func main() {
	treeData := []int64{10, 11, 9, 7, 12, 15, 0} // If you need to create a nil, then send 0
	nodeToBeInserted := 8
	headNode := treeformation.CreateTree(treeData)
	headNode = insertNodeAtFirstVacantPos(headNode, nodeToBeInserted)
	types.InOrderTraversal(headNode)
}
