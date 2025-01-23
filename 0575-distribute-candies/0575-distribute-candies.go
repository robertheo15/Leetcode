func distributeCandies(candyType []int) int {
	result := len(candyType) / 2
	candyMap := make(map[int]bool, len(candyType))

	for _, candy := range candyType {
		candyMap[candy] = true

        //optimize 1
		if len(candyMap) > result {
			return result
		}
	}

    //optimize 2
	// if len(candyMap) < result {
	// 	result = len(candyMap)
	// }

	return min(len(candyMap), len(candyType)/2)
}