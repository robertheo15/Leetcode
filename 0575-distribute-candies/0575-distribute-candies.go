func distributeCandies(candyType []int) int {
	result := len(candyType) / 2
	candyMap := make(map[int]int)

	for _, candy := range candyType {
		candyMap[candy]++
	}
    
	if len(candyMap) < result {
		result = len(candyMap)
	}

	return result
}