func lengthOfLongestSubstring(s string) int {
	lenOfS := len(s)
	maxLength := 0
	lastIndex := make(map[int]int)

	start := 0
	for end := 0; end < lenOfS; end++ {
		currentChar := s[end]
		if lastIndex[int(currentChar)] > start {
			start = lastIndex[int(currentChar)]
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		lastIndex[int(currentChar)] = end + 1
	}

	return maxLength
}