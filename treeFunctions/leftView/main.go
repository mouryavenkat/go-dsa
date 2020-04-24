package main

import (
	"fmt"
	"go-dsa/treeformation"
)

func leftView(headNode *treeformation.Node) {
	count := 1
	nodeMap := map[int][]*treeformation.Node{}
	nodeMap[count] = []*treeformation.Node{headNode}
	for len(nodeMap[count]) != 0 {
		for _, node := range nodeMap[count] {
			if node.Left != nil {
				nodeMap[count+1] = append(nodeMap[count+1], node.Left)
			}
			if node.Right != nil {
				nodeMap[count+1] = append(nodeMap[count+1], node.Right)
			}
		}
		count += 1
	}
	for _, element := range nodeMap {
		fmt.Println(element[0].Data)
	}
}

func main() {
	treeData := []int64{1, 2, 3, 4, 5, 6, 7, 0, 8} // If you need to create a nil, then send 0
	headNode := treeformation.CreateTree(treeData)
	leftView(headNode)
}
