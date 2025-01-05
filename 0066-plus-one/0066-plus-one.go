func plusOne(digits []int) []int {
	ans := digits
	c := 0
    
	if ans[len(ans)-1] != 9 {
		ans[len(ans)-1] += 1
	} else {
		for i := len(ans) - 1; i >= 0; i-- {
			if ans[i] == 9 {
				ans[i] = 0
				c += 1
			} else {
				break
			}
		}
		if c == len(ans) {
			ans[0] = 1
			ans = append(ans, 0)
		} else {
			ans[len(ans)-(c+1)] += 1
		}
	}

	return ans
}