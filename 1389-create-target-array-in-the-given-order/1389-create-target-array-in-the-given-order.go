func createTargetArray(nums []int, index []int) []int {
	results := make([]int, 0, len(nums))

	for i, idx := range index {
		results = append(results[:idx], append([]int{nums[i]}, results[idx:]...)...)
	}

	return results
}