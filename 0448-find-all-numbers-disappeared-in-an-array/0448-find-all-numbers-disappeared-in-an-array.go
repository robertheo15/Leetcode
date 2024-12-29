func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}

	results := make([]int, length)

	for _, num := range nums {
		results[num-1] = num
	}

	j := 0

	for i := 0; i < length; i++ {
		if results[i] == 0 {
			results[j] = i + 1
			j++
		}
	}

	return results[:j]
}