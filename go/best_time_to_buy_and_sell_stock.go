func maxProfit(prices []int) int {
	maxVal := 0
	minAmount := prices[0]

	for i := 1; i < len(prices); i++ {
		if minAmount > prices[i] {
			minAmount = prices[i]
			continue
		}
		if maxVal < prices[i]-minAmount {
			maxVal = prices[i] - minAmount
		}

	}

	return maxVal
}