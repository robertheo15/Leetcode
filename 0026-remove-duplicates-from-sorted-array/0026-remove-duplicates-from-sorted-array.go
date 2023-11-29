func removeDuplicates(nums []int) int {
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[l] = nums[i]
			l++
		}
	}

	return l
}