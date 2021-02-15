package main

import (
	"go-dsa/treeTraversals/traversals/types"
	"go-dsa/treeformation"
)

func convertBSTToSumTree(root *treeformation.Node, data int64) int64 {
	if root == nil {
		return data
	}
	right := convertBSTToSumTree(root.Right, data)
	root.Data = root.Data + right
	left := convertBSTToSumTree(root.Left, root.Data)
	return left
}

func main() {
	treeData := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 0, 0, 17, 18} // If you need to create a nil, then send 0
	headNode := treeformation.CreateTree(treeData)
	convertBSTToSumTree(headNode, 0)
	types.PreOrderTraversal(headNode)
}
