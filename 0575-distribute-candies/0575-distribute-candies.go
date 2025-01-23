func distributeCandies(candyType []int) int {
	// result := len(candyType) / 2
	candyMap := make(map[int]bool, len(candyType))

	for _, candy := range candyType {
		candyMap[candy] = true
	}

	// return min(len(candyMap), result)
    if (len(candyType)/2) > len(candyMap) {
        return len(candyMap)
    } else {
        return len(candyType)/2
    }

    return 1
}