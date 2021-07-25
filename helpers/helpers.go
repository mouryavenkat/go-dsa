package helpers

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


func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min(a,b int) int {
	if a < b {
		return a
	}
	return b
}