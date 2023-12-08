func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	memo := make([]int, 26)

	for i := 0; i < len(s); i++ {
		memo[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		if memo[t[i]-'a'] == 0 {
			return false
		}
		memo[t[i]-'a']--
	}

	return true
}