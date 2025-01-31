func smallerNumbersThanCurrent(nums []int) []int {
	results := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				results[i] += 1
			}
		}
	}

	return results
}