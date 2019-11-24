package main

import (
	"fmt"
	"go-dsa/treeformation"
)

func spiralView(headNode *treeformation.Node) {
	count := 1
	nodeMap := map[int][]*treeformation.Node{}
	nodeMap[count] = []*treeformation.Node{headNode}
	for len(nodeMap[count]) != 0 {
		if count%2 != 0 {
			for i := len(nodeMap[count]) - 1; i >= 0; i-- {
				if nodeMap[count][i].Left != nil {
					nodeMap[count+1] = append(nodeMap[count+1], nodeMap[count][i].Left)
				}
				if nodeMap[count][i].Right != nil {
					nodeMap[count+1] = append(nodeMap[count+1], nodeMap[count][i].Right)
				}
			}
		} else {
			for i := len(nodeMap[count]) - 1; i >= 0; i-- {
				if nodeMap[count][i].Right != nil {
					nodeMap[count+1] = append(nodeMap[count+1], nodeMap[count][i].Right)
				}
				if nodeMap[count][i].Left != nil {
					nodeMap[count+1] = append(nodeMap[count+1], nodeMap[count][i].Left)
				}
			}
		}
		count += 1
	}
	for level, nodeData := range nodeMap {
		fmt.Println(fmt.Sprintf("-----------Level %v ---------", level))
		for _, node := range nodeData {
			fmt.Print(fmt.Sprintf(" %v ", node.Data))
		}
		fmt.Println()
	}
}

func main() {
	treeData := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 0, 0, 17, 18} // If you need to create a nil, then send 0
	headNode := treeformation.CreateTree(treeData)
	spiralView(headNode)
}
