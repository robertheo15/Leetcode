func distributeCandies(candyType []int) int {
	result := len(candyType) / 2
	candyMap := make(map[int]bool, len(candyType))

	for _, candy := range candyType {
		candyMap[candy] = true

		if len(candyMap) > result {
			return result
		}
	}

	return min(len(candyMap), result)
}