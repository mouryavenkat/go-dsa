package main

import "strings"

/*

	You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning, you are standing at index 0, which is equal to '0'. You can move from index i to index j if the following conditions are fulfilled:

	i + minJump <= j <= min(i + maxJump, s.length - 1), and
	s[j] == '0'.
	Return true if you can reach index s.length - 1 in s, or false otherwise.


	Example 1:

	Input: s = "011010", minJump = 2, maxJump = 3
	Output: true
	Explanation:
	In the first step, move from index 0 to index 3.
	In the second step, move from index 3 to index 5.
	Example 2:

	Input: s = "01101110", minJump = 2, maxJump = 3
	Output: false
 */


func canReach(s string, minJump int, maxJump int) bool {
	sSplit := strings.Split(s, "")
	if sSplit[0] != "0" || sSplit[len(s)-1] != "0" {
		return false
	}
	stack := []int{0}

	lastProcessedMax := minJump
	for len(stack) != 0 {
		poppedElement := stack[0]
		if len(stack) > 1 {
			stack = stack[1:]
		} else {
			stack = []int{}
		}
		if poppedElement+minJump > len(s)-1 {
			continue
		}
		max := poppedElement + maxJump
		if poppedElement+maxJump > len(s)-1 {
			max = len(s) - 1
		}
		min := poppedElement + minJump
		if poppedElement+minJump < lastProcessedMax {
			min = lastProcessedMax
		}
		for i := min; i <= max; i++ {
			if sSplit[i] == "0" {
				if i == len(s)-1 {
					return true
				}
				stack = append(stack, i)
			}
		}
		lastProcessedMax = max + 1
	}
	return false
}



