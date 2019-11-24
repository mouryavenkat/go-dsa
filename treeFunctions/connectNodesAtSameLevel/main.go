package main

import (
	"fmt"
	"go-dsa/treeformation"
)

func inorderTraversal(headNode *treeformation.NodeWihNextRight) {
	if headNode == nil {
		return
	}
	inorderTraversal(headNode.Left)
	fmt.Println(headNode.Data)
	inorderTraversal(headNode.Right)
}

func connectNodes(headNode *treeformation.NodeWihNextRight) {
	count := 0
	nodeMap := map[int][]*treeformation.NodeWihNextRight{}
	nodeMap[0] = []*treeformation.NodeWihNextRight{headNode}
	for len(nodeMap[count]) != 0 {
		nodeMapCount := len(nodeMap[count])
		for i := 0; i < nodeMapCount; i++ {
			if nodeMap[count][0].Left != nil {
				nodeMap[count+1] = append(nodeMap[count+1], nodeMap[count][0].Left)
			}
			if nodeMap[count][0].Right != nil {
				nodeMap[count+1] = append(nodeMap[count+1], nodeMap[count][0].Right)
			}
			firstNode := nodeMap[count][0]
			nodeMap[count] = nodeMap[count][1:]
			if len(nodeMap[count]) > 0 {
				firstNode.NextRight = nodeMap[count][0]
			}
		}
		count += 1
	}
}

func printConnectedNodes(headNode *treeformation.NodeWihNextRight) {
	stack := []*treeformation.NodeWihNextRight{headNode}
	for len(stack) != 0 {
		if stack[0].Left != nil {
			stack = append(stack, stack[0].Left)
		}
		if stack[0].Right != nil {
			stack = append(stack, stack[0].Right)
		}
		nodeToDelete := stack[0]
		stack = stack[1:]
		if nodeToDelete.NextRight == nil {
			fmt.Println("Data -> ", nodeToDelete.Data, ". Next Node is nil")
		} else {
			fmt.Println("Data -> ", nodeToDelete.Data, ". Next Node is ", nodeToDelete.NextRight.Data)
		}
	}
}

func main() {
	treeNodes := []int64{10, 3, 5, 4, 1, 0, 2}
	headNode := treeformation.CreateTreeNextRight(treeNodes)
	connectNodes(headNode)
	printConnectedNodes(headNode)
}
