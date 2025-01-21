func maximumWealth(accounts [][]int) int {
	result := 0
	for _, account := range accounts {
		tempSum := 0
		for _, value := range account {
			tempSum += value

			if tempSum > result {
				result = tempSum
			}
		}
	}

	return result
}