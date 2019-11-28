package main

import (
	"go-dsa/treeformation"
	"log"
)

func inorder(headNode *treeformation.Node, nodes *[]int64) {
	if headNode == nil {
		return
	}
	inorder(headNode.Left, nodes)
	*nodes = append(*nodes, headNode.Data)
	inorder(headNode.Right, nodes)
}

func checkIsBST(headNode *treeformation.Node) bool {
	nodes := []int64{}
	inorder(headNode, &nodes)
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i] > nodes[i+1] {
			return false
		}
	}
	return true
}

func main() {
	treeNodes := []int64{5, 3, 7, 1, 4, 6, 9}
	headNode := treeformation.CreateTree(treeNodes)
	if !checkIsBST(headNode) {
		log.Println("Not a BST")
		return
	}
	log.Println("Hurray, its a BST")
}
