package searches

import (
	"go-dsa/collections"
)

type search struct {
}

func NewBinarySearch() collections.BinarySearch {
	return &search{}
}

func (s search) BinarySearchClosest(elements []int, elementToSearch int) (index int) {
	low := 0
	high := len(elements) - 1
	if elementToSearch < elements[0] {
		return 0
	}
	if elementToSearch > elements[high] {
		return high
	}
	for low <= high {
		mid := (low + high) / 2
		if elementToSearch == elements[mid] {
			return mid
		}
		if elementToSearch > elements[mid] {
			low = mid + 1
		}
		if elementToSearch < elements[mid] {
			high = mid - 1
		}
	}
	if (elements[low] - elementToSearch) < (elementToSearch - elements[high]) {
		return low
	}
	return high
}
