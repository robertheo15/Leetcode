func decompressRLElist(nums []int) []int {
	results := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
        freq := nums[i]
        val:=nums[i+1]
		for count := 0; count < freq; count++ {
			results = append(results, val)
		}
	}

	return results
}