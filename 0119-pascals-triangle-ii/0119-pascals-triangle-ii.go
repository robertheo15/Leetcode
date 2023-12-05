func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1

	if rowIndex == 0 {
		return res
	}

	prev := 1
	for i := 1; i < rowIndex+1; i++ {
		curr := prev * (rowIndex - i + 1) / i
		prev = curr
		res[i] = curr
	}

	return res
}