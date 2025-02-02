func mostWordsFound(sentences []string) int {
	counter := 0

	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		if len(words) > counter {
			counter = len(words)
		}
	}
	
	return counter
}