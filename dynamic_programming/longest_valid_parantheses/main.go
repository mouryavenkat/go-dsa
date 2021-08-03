package longest_valid_parantheses

import "strings"

func getTopElement(stack []int, sSplit []string) (string, int) {
	return sSplit[stack[len(stack)-1]], stack[len(stack)-1]
}

func pop(stack []int) []int {
	if len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}
	return stack
}

/*
	Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

	Example 1:

	Input: s = "(()"
	Output: 2
	Explanation: The longest valid parentheses substring is "()".
	Example 2:

	Input: s = ")()())"
	Output: 4
	Explanation: The longest valid parentheses substring is "()()".
	Example 3:

	Input: s = ""
	Output: 0


*/
func longestValidParentheses(s string) int {
	bestAtIndex := make([]int, len(s))
	var stack []int
	sSplit := strings.Split(s, "")
	firstValidIndex := -1
	for index, char := range sSplit {
		if char == "(" {
			if firstValidIndex == -1 {
				firstValidIndex = index
			}
			stack = append(stack, index)
		} else if char == ")" {
			if len(stack) == 0 {
				firstValidIndex = -1
				continue
			}
			topElement, topElementIndex := getTopElement(stack, sSplit)
			if topElement == "(" {
				bestAtIndex[topElementIndex] = bestAtIndex[topElementIndex] + 2
				stack = pop(stack)
				if len(stack) > 0 {
					_, topElementIndexPrev := getTopElement(stack, sSplit)
					bestAtIndex[topElementIndexPrev] = bestAtIndex[topElementIndexPrev] + bestAtIndex[topElementIndex]
				} else if topElementIndex != firstValidIndex && len(stack) == 0 {
					bestAtIndex[firstValidIndex] = bestAtIndex[firstValidIndex] + bestAtIndex[topElementIndex]
				}
			}
		}
	}
	max := 0
	for _, element := range bestAtIndex {
		if max < element {
			max = element
		}
	}
	return max
}
