package main

import "fmt"

func StringPermutations(inputStr string) (permutations []string) {
	for i := 0; i < len(inputStr); i++ {
		inputStr = inputStr[1:] + inputStr[0:1]
		if len(inputStr) <= 2 {
			permutations = append(permutations, inputStr)
			continue
		}
		leftPermutations := StringPermutations(inputStr[0 : len(inputStr)-1])
		for _, leftPermutation := range leftPermutations {
			permutations = append(permutations, leftPermutation+inputStr[len(inputStr)-1:])
		}
	}
	return
}

func main() {
	fmt.Println(len(StringPermutations("ABCDE")))
}
