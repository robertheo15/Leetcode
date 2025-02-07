func decompressRLElist(nums []int) []int {
	results := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
		for count := 0; count < nums[i]; count++ {
			results = append(results, nums[i+1])
		}
	}

	return results
}