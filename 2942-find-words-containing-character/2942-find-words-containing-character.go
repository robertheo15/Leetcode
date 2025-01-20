func findWordsContaining(words []string, x byte) []int {
	var nums []int

	for i, word := range words {
        if strings.ContainsRune(word, rune(x)) {
			nums = append(nums, i)
		}
	}

	return nums
}

// func findWordsContaining(words []string, x byte) []int {
// 	var nums []int

// 	for i, word := range words {
// 		for _, c := range word {
// 			if byte(c) == x {
// 				nums = append(nums, i)
// 				break
// 			}
// 		}
// 	}

// 	return nums
// }