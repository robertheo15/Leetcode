func finalPrices(prices []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] && j == len(prices)-1 {
				result = append(result, prices[i])
				break
			}

			if prices[j] <= prices[i] && j > i {
				result = append(result, prices[i]-prices[j])
				break
			}
		}
	}
    
	result = append(result, prices[len(prices)-1])

	return result
}