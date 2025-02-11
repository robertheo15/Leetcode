func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	avg := make([]float64, 0)

	for i := 0; i < len(nums)/2; i++ {
		left, right := nums[i], nums[len(nums)-i-1]
		avg = append(avg, (float64(left)+float64(right))/2)
	}

	result := slices.Min(avg)

	return result
}