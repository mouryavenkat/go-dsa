package main

import (
	"sort"
)

/*
	There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0], and the cost of flying the i-th person to city B is costs[i][1].

	Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.

	Example 1:

	Input: [[10,20],[30,200],[400,50],[30,20]]
	Output: 110
	Explanation:
	The first person goes to city A for a cost of 10.
	The second person goes to city A for a cost of 30.
	The third person goes to city B for a cost of 50.
	The fourth person goes to city B for a cost of 20.

	The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city
*/

func twoCitySchedCost(costs [][]int) int {
	costsLen := len(costs)
	sort.SliceStable(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < costs[j][0]-costs[j][1]
	})
	sum := 0
	for i := 0; i < costsLen/2; i++ {
		sum = sum + costs[i][0]
	}
	for i := costsLen / 2; i < len(costs); i++ {
		sum = sum + costs[i][1]
	}
	return sum
}

func main() {
	costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	twoCitySchedCost(costs)
}
