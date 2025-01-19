func shuffle(nums []int, n int) []int {
	result := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[i+n])
	}
	return result
}