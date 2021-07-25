package main

import "fmt"

func main() {
	arr := []int{4, 6, 3, 9, 5, 1 ,7}
	nextGreatest := make([]int, len(arr))
	var stack []int
	for index, element := range arr {
		if len(stack) == 0 {
			stack = append(stack, index)
			continue
		}
		if element > arr[stack[len(stack)-1]] {
			for len(stack) > 0 && element > arr[stack[len(stack)-1]] {
				stackElement := stack[len(stack)-1]
				stack = stack[0:len(stack)-1]
				nextGreatest[stackElement] = element
			}
		}
		stack = append(stack, index)
	}
	for len(stack) > 0 {
		nextGreatest[stack[len(stack)-1]] = -1
		stack = stack[0:len(stack)-1]
	}
	fmt.Println(nextGreatest)
}