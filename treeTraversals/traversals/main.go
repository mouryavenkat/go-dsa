package main

import (
	"go-dsa/treeTraversals/traversals/types"
	"go-dsa/treeformation"
)

func main() {
	treeData := []int64{10, 11, 9, 7, 0, 15, 8} // If you need to create a nil, then send 0
	headNode := treeformation.CreateTree(treeData)
	types.InOrderTraversal(headNode)
	types.PreOrderTraversal(headNode)
	types.PostOrderTraversal(headNode)
}
