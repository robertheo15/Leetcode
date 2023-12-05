func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if profit := prices[i] - minPrice; profit > maxProfit {
			// If diff of current stock with minPrice is greater
			// update the profit
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
