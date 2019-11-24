package main

import (
	"go-dsa/treeformation"
	"fmt"
	"sort"
)

func bottomView(headNode *treeformation.Node, count int, nodeMap map[int][]*treeformation.Node, ltype string) {
	if headNode == nil {
		return
	}
	if headNode.Left != nil {
		if ltype == "left" {
			nodeMap[count-1] = append([]*treeformation.Node{headNode.Left}, nodeMap[count-1]...)
		} else {
			nodeMap[count-1] = append(nodeMap[count-1], headNode.Left)
		}
	}
	bottomView(headNode.Left, count-1, nodeMap, "left")
	if headNode.Right != nil {
		if ltype == "left" {
			nodeMap[count+1] = append([]*treeformation.Node{headNode.Right}, nodeMap[count+1]...)
		} else {
			nodeMap[count+1] = append(nodeMap[count+1], headNode.Right)
		}
	}
	bottomView(headNode.Right, count+1, nodeMap, "left")
}

func main() {
	treeData := []int64{20, 8, 22, 5, 3, 4, 25, 6, 12, 10, 14} // If you need to create a nil, then send 0
	headNode := treeformation.CreateTree(treeData)
	nodeMap := map[int][]*treeformation.Node{}
	nodeMap[0] = []*treeformation.Node{headNode}
	bottomView(headNode, 0, nodeMap, "")
	keys := []int{}
	for index := range nodeMap {
		keys = append(keys, index)
	}
	sort.Ints(keys)

	for _, index := range keys {
		fmt.Println(nodeMap[index][0].Data)
	}
}
