func reversePrefix(word string, ch byte) string {
	wordToByte := []byte(word)
	tempByte := make([]byte, len(wordToByte))

	for i, char := range wordToByte {
		if char == ch {
			tempByte = append(tempByte[0:i+1], wordToByte[i+1:]...)
			fmt.Println(string(tempByte))
			fmt.Println(string(wordToByte[:i+1]))

			lenTemp := len(wordToByte[:i+1]) - 1
			for j := 0; j <= lenTemp; j++ {
				tempByte[j] = wordToByte[lenTemp-j]
			}

			return string(tempByte)
		} 
	}
    
	return word
}
