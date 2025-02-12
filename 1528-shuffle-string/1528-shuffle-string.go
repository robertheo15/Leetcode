func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))

	for i, num := range indices {
		result[num] = s[i]
	}

	return string(result)
}