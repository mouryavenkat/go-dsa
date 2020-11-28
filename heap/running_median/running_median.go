package main

import (
	"bufio"
	"fmt"
	"go-dsa/heap"
	"os"
	"strconv"
	"strings"
)

func main() {
	leftHeapObject := heap.NewHeap()
	rightHeapObject := heap.NewHeap()
	var leftNodes, rightNodes int
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		intInput, err := strconv.ParseInt(strings.TrimSuffix(input, "\n"), 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		if leftNodes == 0 || leftHeapObject.HeapNodes[0].Data >= int(intInput) {
			leftHeapObject.Insert(&heap.Node{
				Val:  int(intInput),
				Data: int(intInput),
			}, func(a int, b int) bool {
				return a <= b
			})
			leftNodes++
		} else {
			rightHeapObject.Insert(&heap.Node{
				Val:  int(intInput),
				Data: int(intInput),
			}, func(a int, b int) bool {
				return a >= b
			})
			rightNodes++
		}
		if leftNodes-rightNodes > 1 {
			deletedNode := leftHeapObject.Delete(func(a int, b int) bool {
				return a < b
			})
			leftNodes--
			rightHeapObject.Insert(deletedNode, func(a int, b int) bool {
				return a >= b
			})
			rightNodes++
		} else if rightNodes-leftNodes > 1 {
			deletedNode := rightHeapObject.Delete(func(a int, b int) bool {
				return a > b
			})
			rightNodes--
			leftHeapObject.Insert(deletedNode, func(a int, b int) bool {
				return a <= b
			})
			leftNodes++
		}
		if leftNodes == rightNodes {
			fmt.Println("Median", (leftHeapObject.HeapNodes[0].Data+rightHeapObject.HeapNodes[0].Data)/2)
		} else if leftNodes > rightNodes {
			fmt.Println("Median", leftHeapObject.HeapNodes[0].Data)
		} else {
			fmt.Println("Median", rightHeapObject.HeapNodes[0].Data)
		}

	}
}
