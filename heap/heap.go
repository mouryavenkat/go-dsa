package main

import (
	"fmt"
)

type node struct {
	vertex int
	cost   int
	stops  int
}

type shortestPath struct {
	nodes          []*node
	count          int
}

func New() *shortestPath {
	return &shortestPath{
		nodes:          []*node{},
		count:          0,
	}
}

func (s *shortestPath) print() {
	arr := [][]int{}
	for _, node := range s.nodes {
		arr = append(arr, []int{node.vertex, int(node.cost)})
	}
	fmt.Println(arr)
}

func (s *shortestPath) heapify() {
	nodeCount := s.count
	for parentIndex := nodeCount - 1; parentIndex >= 0; parentIndex-- {
		for parentIndex < nodeCount {
			leftChild := 2*parentIndex + 1
			rightChild := 2 * (parentIndex + 1)
			if leftChild >= nodeCount && rightChild >= nodeCount {
				break
			}
			// best child mean least cost
			var bestChild int
			if (rightChild >= nodeCount) || (s.nodes[leftChild].cost < s.nodes[rightChild].cost) {
				bestChild = leftChild
			} else {
				bestChild = rightChild
			}
			if s.nodes[parentIndex].cost <= s.nodes[bestChild].cost {
				break
			}
			s.nodes[parentIndex], s.nodes[bestChild] = s.nodes[bestChild], s.nodes[parentIndex]
			parentIndex = bestChild
		}
	}
}

// Min Heap
func (s *shortestPath) insert(newNode *node) {

	s.nodes = append(s.nodes, newNode)
	s.count++
	nodeCount := s.count
	parentIndex := ((nodeCount + 1) / 2) - 1
	if parentIndex < 0 {
		return
	}
	for parentIndex >= 0 && s.nodes[nodeCount-1].cost < s.nodes[parentIndex].cost {
		s.nodes[nodeCount-1], s.nodes[parentIndex] = s.nodes[parentIndex], s.nodes[nodeCount-1]
		nodeCount = parentIndex + 1
		parentIndex = (nodeCount / 2) - 1
	}

}

func (s *shortestPath) delete() *node {
	nodeCount := s.count
	lastNode := s.nodes[nodeCount-1]
	firstNode := s.nodes[0]
	s.nodes = s.nodes[0 : nodeCount-1]
	s.count--
	if s.count == 0 {
		return firstNode
	}
	s.nodes[0] = lastNode
	nodeCount--
	parentIndex := 0
	for parentIndex < nodeCount {
		leftChild := 2*parentIndex + 1
		rightChild := 2 * (parentIndex + 1)
		if leftChild >= nodeCount && rightChild >= nodeCount {
			return firstNode
		}
		// best child mean least cost
		var bestChild int
		if (rightChild >= nodeCount) || (s.nodes[leftChild].cost < s.nodes[rightChild].cost) {
			bestChild = leftChild
		} else {
			bestChild = rightChild
		}
		if s.nodes[parentIndex].cost > s.nodes[bestChild].cost {
			s.nodes[parentIndex], s.nodes[bestChild] = s.nodes[bestChild], s.nodes[parentIndex]
		}
		parentIndex = bestChild
	}
	return firstNode
}


