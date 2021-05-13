package palindromic_substrings

func countSubstrings(s string) int {
	sSplit := strings.Split(s, "")
	prevData := make([][]bool, len(sSplit))
	var trueCount int
	for i := 0; i < len(prevData); i++ {
		prevData[i] = make([]bool, len(sSplit))
		prevData[i][i] = true
		trueCount++
	}
	for i := 1; i < len(sSplit); i++ {
		for j := 0; j < len(sSplit)-i; j++ {
			startIndex := j
			endIndex := j + i
			if startIndex+1 >= endIndex {
				if sSplit[startIndex] == sSplit[endIndex] {
					prevData[startIndex][endIndex] = true
					trueCount++
				}
				continue
			}
			if prevData[startIndex+1][endIndex-1] == true && sSplit[startIndex] == sSplit[endIndex] {
				prevData[startIndex][endIndex] = true
				trueCount++
			}
		}
	}
	return trueCount
}
