func leftRightDifference(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, 0

	for _, num := range nums {
		right += num
	}

	for i, num := range nums {
		right -= num
		result[i] = abs(left - right)
		left += num
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}