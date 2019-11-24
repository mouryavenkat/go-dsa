package types

import (
	"fmt"
	"go-dsa/treeformation"
)

func InOrderTraversal(headNode *treeformation.Node) {
	if headNode == nil {
		return
	}
	InOrderTraversal(headNode.Left)
	fmt.Println(headNode.Data)
	InOrderTraversal(headNode.Right)
}

func PreOrderTraversal(headNode *treeformation.Node) {
	if headNode == nil {
		return
	}
	fmt.Println(headNode.Data)
	InOrderTraversal(headNode.Left)
	InOrderTraversal(headNode.Right)
}

func PostOrderTraversal(headNode *treeformation.Node) {
	if headNode == nil {
		return
	}
	InOrderTraversal(headNode.Left)
	InOrderTraversal(headNode.Right)
	fmt.Println(headNode.Data)
}