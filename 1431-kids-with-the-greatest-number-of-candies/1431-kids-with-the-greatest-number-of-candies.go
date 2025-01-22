import "slices"
func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := slices.Max(candies)
	// for i := 0; i < len(candies); i++ {
	// 	if temp < candies[i] {
	// 		temp = candies[i]
	// 	}
	// }

	for i, candy := range candies {
		if candy+extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}