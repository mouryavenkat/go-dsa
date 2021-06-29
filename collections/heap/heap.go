package heap

import (
	"fmt"
	"go-dsa/collections"
)

type node struct {
	val  int
	data int
}

type heapObject struct {
	heapNodes []*node
	count     int
}

func NewHeap() collections.Heap {
	return &heapObject{
		heapNodes: []*node{},
		count:     0,
	}
}

func (s *heapObject) Print() {
	for _, nodes := range s.heapNodes {
		fmt.Print(fmt.Sprintf("%d\t", nodes.data))
	}
	fmt.Println("\n--------")
}

func (s *heapObject) Insert(val, data int, sortFunction func(a int, b int) bool) {
	newNode := &node{
		val:  val,
		data: data,
	}
	s.heapNodes = append(s.heapNodes, newNode)
	s.count++
	nodeCount := s.count - 1
	parent := (nodeCount - 1) / 2
	for parent >= 0 {
		if sortFunction(s.heapNodes[nodeCount].data, s.heapNodes[parent].data) {
			return
		}
		s.heapNodes[nodeCount], s.heapNodes[parent] = s.heapNodes[parent], s.heapNodes[nodeCount]
		nodeCount = parent
		parent = (nodeCount - 1) / 2
	}
}

func (s *heapObject) heapify(parentIndex int, sortFunction func(a int, b int) bool) {
	leftChild := 2*parentIndex + 1
	rightChild := 2 * (parentIndex + 1)
	bestIndex := parentIndex
	if leftChild < s.count && sortFunction(s.heapNodes[parentIndex].data, s.heapNodes[leftChild].data) {
		bestIndex = leftChild
	}
	if rightChild < s.count && sortFunction(s.heapNodes[bestIndex].data, s.heapNodes[rightChild].data) {
		bestIndex = rightChild
	}
	if bestIndex == parentIndex {
		return
	}
	s.heapNodes[parentIndex], s.heapNodes[bestIndex] = s.heapNodes[bestIndex], s.heapNodes[parentIndex]
	s.heapify(bestIndex, sortFunction)
}

func (s *heapObject) Delete(sortFunction func(a int, b int) bool) (val, data int, exists bool) {
	if s.count == 0 {
		return -1, -1, false
	}
	firstNode := s.heapNodes[0]
	lastNode := s.heapNodes[s.count-1]
	s.heapNodes = s.heapNodes[0 : s.count-1]
	s.count--
	if s.count == 0 {
		return firstNode.val, firstNode.data, true
	}
	s.heapNodes[0] = lastNode
	s.heapify(0, sortFunction)
	return firstNode.val, firstNode.data, true
}

func (s *heapObject) GetDataAtIndex(index int) int {
	return s.heapNodes[index].data
}
