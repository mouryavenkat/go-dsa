package main

/*
	Say you have an array for which the ith element is the price of a given stock on day i.

	If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

	Note that you cannot sell a stock before you buy one.

	Example 1:

	Input: [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
				 Not 7-1 = 6, as selling price needs to be larger than buying price.
	Example 2:

	Input: [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

func bestOfTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	priceLen := len(prices)
	if priceLen == 0 {
		return 0
	}
	bestPrices := make([]int, priceLen)
	bestPrices[priceLen-1] = -1
	for i := priceLen - 2; i >= 0; i-- {
		nextBest := bestOfTwo(prices[i+1], bestPrices[i+1])
		if nextBest > prices[i] {
			bestPrices[i] = nextBest
		} else {
			bestPrices[i] = -1
		}
	}
	maxDiff := 0
	for i := 0; i < priceLen; i++ {
		if bestPrices[i] > 0 && bestPrices[i]-prices[i] > maxDiff {
			maxDiff = bestPrices[i] - prices[i]
		}
	}
	return maxDiff
}
