func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]

	for _, price := range prices {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// If diff of current stock with minPrice is greater
			// update the profit
			maxProfit = profit
		}
	}

	return maxProfit
}