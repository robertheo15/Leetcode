func truncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
    words = words[:k]

	return strings.Join(words, " ")
}