func maxDepth(s string) int {
	counter, maxCounter := 0, 0
	for i := range s {
		if s[i] == '(' {
			counter++
			maxCounter = max(maxCounter, counter)
		} else if s[i] == ')' {
			counter--
			maxCounter = max(maxCounter, counter)
		}
	}

	return maxCounter
}