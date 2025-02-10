func numberGame(nums []int) []int {
	results := make([]int, len(nums))
	sort.Ints(nums)

	for i := 0; i < len(results)-1; i += 2 {
		results[i], results[i+1] = nums[i+1], nums[i]
	}

	return results
}