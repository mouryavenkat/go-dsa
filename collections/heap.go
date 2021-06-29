package collections

type Heap interface {
	Insert(val, data int, sortFunction func(a int, b int) bool)
	Delete(sortFunction func(a int, b int) bool) (val, data int, exists bool)
	GetDataAtIndex(index int) int
}
