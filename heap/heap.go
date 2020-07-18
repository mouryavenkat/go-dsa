package heap

type Node struct {
	Val  int
	Data int
}

type HeapSort struct {
	HeapNodes []*Node
	Count     int
}

func NewHeap() *HeapSort {
	return &HeapSort{
		HeapNodes: []*Node{},
		Count:     0,
	}
}

func (s *HeapSort) Insert(newNode *Node) {
	s.HeapNodes = append(s.HeapNodes, newNode)
	s.Count++
	nodeCount := s.Count - 1
	parent := (nodeCount - 1) / 2
	for parent >= 0 {
		if s.HeapNodes[nodeCount].Data <= s.HeapNodes[parent].Data {
			return
		}
		s.HeapNodes[nodeCount], s.HeapNodes[parent] = s.HeapNodes[parent], s.HeapNodes[nodeCount]
		nodeCount = parent
		parent = (nodeCount - 1) / 2
	}
}

func (s *HeapSort) heapify(parentIndex int) {
	leftChild := 2*parentIndex + 1
	rightChild := 2 * (parentIndex + 1)
	bestIndex := parentIndex
	if leftChild < s.Count && s.HeapNodes[parentIndex].Data < s.HeapNodes[leftChild].Data {
		bestIndex = leftChild
	}
	if rightChild < s.Count && s.HeapNodes[bestIndex].Data < s.HeapNodes[rightChild].Data {
		bestIndex = rightChild
	}
	if bestIndex == parentIndex {
		return
	}
	s.HeapNodes[parentIndex], s.HeapNodes[bestIndex] = s.HeapNodes[bestIndex], s.HeapNodes[parentIndex]
	s.heapify(bestIndex)
}

func (s *HeapSort) Delete() *Node {
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
	s.heapify(0)
	return firstNode
}
