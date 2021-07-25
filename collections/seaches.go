package collections

type BinarySearch interface {
	BinarySearchClosest(elements []int, elementToSearch int) (index int)
}