func reverseString(s []byte) {
	length := len(s)-1

	for i := 0; i <= length; i++ {
        s[i],s[length] = s[length],s[i]
        length--
	}
}