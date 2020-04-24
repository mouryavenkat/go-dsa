package main

import (
	"fmt"
	"go-dsa/treeformation"
)

// One approach would be to get nodes in pre-order and post-order and compare each of them.
func checkMirror(headNode *treeformation.Node) bool {
	leftStack := []*treeformation.Node{headNode.Left}
	rightStack := []*treeformation.Node{headNode.Right}
	leftStackLen := len(leftStack)
	for i := 0; i < leftStackLen; i++ {
		if leftStack[0] == nil {
			if rightStack[0] != nil {
				return false
			}
			leftStack = leftStack[1:]
			rightStack = rightStack[1:]
			continue
		}
		if rightStack[0] == nil {
			if leftStack[0] != nil {
				return false
			}
			leftStack = leftStack[1:]
			rightStack = rightStack[1:]
			continue
		}
		if leftStack[0].Data != rightStack[0].Data {
			return false
		}
		leftStack = append(leftStack, []*treeformation.Node{leftStack[0].Left, leftStack[0].Right}...)
		rightStack = append(rightStack, []*treeformation.Node{rightStack[0].Right, rightStack[0].Left}...)
		leftStack = leftStack[1:]
		rightStack = rightStack[1:]
		leftStackLen = len(leftStack)
	}
	return true
}
func main() {
	treeNodes := []int64{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5}
	headNode := treeformation.CreateTree(treeNodes)
	fmt.Println(checkMirror(headNode))
}
