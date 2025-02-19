func clearDigits(s string) string {
	tempString := make([]rune, 0)
	for _, char := range s {
		if char >= '0' && char <= '9' {
			if len(tempString) > 0 && tempString[len(tempString)-1] >= 'a' && tempString[len(tempString)-1] <= 'z' {
				tempString = tempString[:len(tempString)-1]
			}
		} else {
			tempString = append(tempString, char)
		}
	}
    return string(tempString)
}