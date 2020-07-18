package heap

type Node struct {
	Val  int
	Data int
}

type HeapSort struct {
	heapNodes []*Node
	count     int
}

func NewHeap() *HeapSort {
	return &HeapSort{
		heapNodes: []*Node{},
		count:     0,
	}
}

func (s *HeapSort) Insert(newNode *Node) {
	s.heapNodes = append(s.heapNodes, newNode)
	s.count++
	nodeCount := s.count - 1
	parent := (nodeCount - 1) / 2
	for parent >= 0 {
		if s.heapNodes[nodeCount].Data <= s.heapNodes[parent].Data {
			return
		}
		s.heapNodes[nodeCount], s.heapNodes[parent] = s.heapNodes[parent], s.heapNodes[nodeCount]
		nodeCount = parent
		parent = (nodeCount - 1) / 2
	}
}

func (s *HeapSort) heapify(parentIndex int) {
	leftChild := 2*parentIndex + 1
	rightChild := 2 * (parentIndex + 1)
	bestIndex := parentIndex
	if leftChild < s.count && s.heapNodes[parentIndex].Data < s.heapNodes[leftChild].Data {
		bestIndex = leftChild
	}
	if rightChild < s.count && s.heapNodes[bestIndex].Data < s.heapNodes[rightChild].Data {
		bestIndex = rightChild
	}
	if bestIndex == parentIndex {
		return
	}
	s.heapNodes[parentIndex], s.heapNodes[bestIndex] = s.heapNodes[bestIndex], s.heapNodes[parentIndex]
	s.heapify(bestIndex)
}

func (s *HeapSort) Delete() *Node {
	if s.count == 0 {
		return nil
	}
	firstNode := s.heapNodes[0]
	lastNode := s.heapNodes[s.count-1]
	s.heapNodes = s.heapNodes[0 : s.count-1]
	s.count--
	if s.count == 0 {
		return firstNode
	}
	s.heapNodes[0] = lastNode
	s.heapify(0)
	return firstNode
}
