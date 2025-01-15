func sortedInts(nums []int) []int {
	copyNums := append([]int{}, nums...)
	sort.Ints(copyNums)
	return copyNums
}

func heightChecker(heights []int) int {
	expected := sortedInts(heights)
	counter := 0
	for i, height := range heights {
		if height != expected[i] {
			counter++
		}
	}

	return counter
}
