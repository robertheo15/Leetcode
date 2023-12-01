func plusOne(digits []int) []int {
	ans := digits
	c := 0
	//If the last digit is not a nine, we don't have to worry about carrying so we just return the array with the last digit being increased by 1
	if ans[len(ans)-1] != 9 {
		ans[len(ans)-1] += 1
	} else {
		//If the last digit is a nine then we count (c) to see how many 9's there are. We iterate backwards through the array
		for i := len(ans) - 1; i >= 0; i-- {
			if ans[i] == 9 {
				//if we find a nine we change it to a zero and increase the amount of digits we carry over (c)
				ans[i] = 0
				c += 1
				//If we find a number that isn't a nine then we break the loop
			} else {
				break
			}
		}
		//If all of the digits are nine's then we make the first digit a one and add a new zero to the end of the array (ex. 99 becomes 100)
		if c == len(ans) {
			ans[0] = 1
			ans = append(ans, 0)
		} else {
			//Otherwise just add one to whichever digit is a non-nine int.
			ans[len(ans)-(c+1)] += 1
		}
	}
	//Return answer
	return ans
}
