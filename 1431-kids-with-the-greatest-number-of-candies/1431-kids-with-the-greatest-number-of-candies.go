import "slices"
func kidsWithCandies(candies []int, extraCandies int) []bool {
	results := make([]bool, len(candies))
    max := slices.Max(candies)

	for i, candy := range candies {
		if candy+extraCandies >= max {
			results[i] = true
		} else {
			results[i] = false
		}
	}

	return results
}