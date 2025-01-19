func shuffle(nums []int, n int) []int {
	results := make([]int, 2*n)
	for i := 0; i < n; i++  {
        results[2*i] = nums[i]
        results[2*i+1] = nums[n+i]
	}

	return results
}