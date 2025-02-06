func subarraySum(nums []int) int {
	subSum := 0
	for index, num := range nums {
		start := max(0, index-num)
		for i := start; i <= index; i++ {
			subSum += nums[i]
		}
	}

	return subSum
}