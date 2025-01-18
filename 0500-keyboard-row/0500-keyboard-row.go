func findWords(words []string) []string {
	rowMap := populateMap()
	result := make([]string, 0)

	for _, word := range words {
		wordLower := strings.ToLower(word)
		row := rowMap[rune(wordLower[0])]
		valid := true

		for _, c := range wordLower {
			if rowMap[c] != row {
				valid = false
				break
			}
		}

		if valid {
			result = append(result, word)
		}
	}

	return result
}

func populateMap() map[rune]int {
	rowMap := map[rune]int{}
	for _, c := range "qwertyuiop" {
		rowMap[c] = 1
	}
	for _, c := range "asdfghjkl" {
		rowMap[c] = 2
	}
	for _, c := range "zxcvbnm" {
		rowMap[c] = 3
	}

	return rowMap
}