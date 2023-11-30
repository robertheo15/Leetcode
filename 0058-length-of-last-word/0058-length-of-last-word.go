func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")

	if len(words) == 1 {
		return len(words[0])
	}

	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return len(words[len(words)-1])
}