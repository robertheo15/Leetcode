func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 { // handle only 1 element
		return strs[0]
	}

	// sort them first, the most different one will be in first and last
	sort.Strings(strs)

	// compare first and last
	l := len(strs)
	for i := range strs[0] {
		fmt.Println(strs[0][i])
		fmt.Println(strs[l-1][i])
		if strs[0][i] != strs[l-1][i] {
			return strs[0][:i]
		}
	}
	return strs[0]
}