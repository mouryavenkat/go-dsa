package main

import (
	"bufio"
	"fmt"
	heap2 "go-dsa/collections/heap"
	"os"
	"strconv"
	"strings"
)

func main() {
	leftHeapObject := heap2.NewHeap()
	rightHeapObject := heap2.NewHeap()
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
		if leftNodes == 0 || leftHeapObject.GetDataAtIndex(0) >= int(intInput) {
			leftHeapObject.Insert(int(intInput), int(intInput), func(a int, b int) bool {
				return a <= b
			})
			leftNodes++
		} else {
			rightHeapObject.Insert(int(intInput), int(intInput), func(a int, b int) bool {
				return a >= b
			})
			rightNodes++
		}
		if leftNodes-rightNodes > 1 {
			val, data, _ := leftHeapObject.Delete(func(a int, b int) bool {
				return a < b
			})
			leftNodes--
			rightHeapObject.Insert(val, data, func(a int, b int) bool {
				return a >= b
			})
			rightNodes++
		} else if rightNodes-leftNodes > 1 {
			val, data, _ := rightHeapObject.Delete(func(a int, b int) bool {
				return a > b
			})
			rightNodes--
			leftHeapObject.Insert(val, data, func(a int, b int) bool {
				return a <= b
			})
			leftNodes++
		}
		if leftNodes == rightNodes {
			fmt.Println("Median", (leftHeapObject.GetDataAtIndex(0)+rightHeapObject.GetDataAtIndex(0))/2)
		} else if leftNodes > rightNodes {
			fmt.Println("Median", leftHeapObject.GetDataAtIndex(0))
		} else {
			fmt.Println("Median", rightHeapObject.GetDataAtIndex(0))
		}

	}
}
