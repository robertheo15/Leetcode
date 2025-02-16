func removeOuterParentheses(s string) string {
	result := ""
	counter := 0
	for _, ch := range s {
		if ch == '(' {
            if counter > 0 {
                result += string(ch)
            }
			counter++
		} else if ch == ')' {
            counter--
            if counter > 0 {
                result += string(ch)
            }
        }
	}

	return result
}