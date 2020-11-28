package heap

import "fmt"

type Node struct {
	Val  int
	Data int
}

type HeapObject struct {
	HeapNodes []*Node
	Count     int
}

func NewHeap() *HeapObject {
	return &HeapObject{
		HeapNodes: []*Node{},
		Count:     0,
	}
}

func (s *HeapObject) Print() {
	for _, nodes := range s.HeapNodes {
		fmt.Print(fmt.Sprintf("%d\t", nodes.Data))
	}
	fmt.Println("\n--------")
}

func (s *HeapObject) Insert(newNode *Node, sortFunction func(a int, b int) bool) {
	s.HeapNodes = append(s.HeapNodes, newNode)
	s.Count++
	nodeCount := s.Count - 1
	parent := (nodeCount - 1) / 2
	for parent >= 0 {
		if sortFunction(s.HeapNodes[nodeCount].Data, s.HeapNodes[parent].Data) {
			return
		}
		s.HeapNodes[nodeCount], s.HeapNodes[parent] = s.HeapNodes[parent], s.HeapNodes[nodeCount]
		nodeCount = parent
		parent = (nodeCount - 1) / 2
	}
}

func (s *HeapObject) heapify(parentIndex int, sortFunction func(a int, b int) bool) {
	leftChild := 2*parentIndex + 1
	rightChild := 2 * (parentIndex + 1)
	bestIndex := parentIndex
	if leftChild < s.Count && sortFunction(s.HeapNodes[parentIndex].Data, s.HeapNodes[leftChild].Data) {
		bestIndex = leftChild
	}
	if rightChild < s.Count && sortFunction(s.HeapNodes[bestIndex].Data, s.HeapNodes[rightChild].Data) {
		bestIndex = rightChild
	}
	if bestIndex == parentIndex {
		return
	}
	s.HeapNodes[parentIndex], s.HeapNodes[bestIndex] = s.HeapNodes[bestIndex], s.HeapNodes[parentIndex]
	s.heapify(bestIndex, sortFunction)
}

func (s *HeapObject) Delete(sortFunction func(a int, b int) bool) *Node {
	if s.Count == 0 {
		return nil
	}
	firstNode := s.HeapNodes[0]
	lastNode := s.HeapNodes[s.Count-1]
	s.HeapNodes = s.HeapNodes[0 : s.Count-1]
	s.Count--
	if s.Count == 0 {
		return firstNode
	}
	s.HeapNodes[0] = lastNode
	s.heapify(0, sortFunction)
	return firstNode
}
