func differenceOfSum(nums []int) int {
	sumElement, sumDigit := 0, 0

	for _, num := range nums {
		sumElement += num
		digit := num % 10
		digit += num / 10
		sumDigit = sumDigit + digit
	}

	return sumElement - sumDigit
}