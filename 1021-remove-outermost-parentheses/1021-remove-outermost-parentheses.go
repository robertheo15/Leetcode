// func removeOuterParentheses(s string) string {
// 	result := ""
// 	counter := 0
// 	for _, ch := range s {
// 		if ch == '(' {
//             if counter > 0 {
//                 result += string(ch)
//             }
// 			counter++
// 		} else if ch == ')' {
//             counter--
//             if counter > 0 {
//                 result += string(ch)
//             }
//         }
// 	}

// 	return result
// }
func removeOuterParentheses(s string) string {
    counter, result := 0, strings.Builder{}
    for i := range s {
        if s[i] == '(' {
            counter++
            if counter != 1 {
                result.WriteByte(s[i])
            }
        } else {
            counter--
            if counter > 0 {
                result.WriteByte(s[i])
            }
        }
    }

    return result.String()
}