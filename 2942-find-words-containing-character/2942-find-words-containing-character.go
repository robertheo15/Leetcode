func findWordsContaining(words []string, x byte) []int {
	nums := make([]int, 0)

	for i, word := range words {
        if strings.ContainsRune(word, rune(x)) {
			nums = append(nums, i)
		}
	}

	return nums
}