func searchInsert(nums []int, target int) int {

	return binarySearch(0, len(nums)-1, nums, target)
}

func binarySearch(a int, b int, nums []int, target int) int {

	c := (b + a) / 2

	if nums[c] == target {
		return c
	}

	if target < nums[c] {
		if a == c {
			return a
		} else {
			return binarySearch(a, c-1, nums, target)
		}
	}

	if b == c {
		return b + 1
	} else {
		return binarySearch(c+1, b, nums, target)
	}
}