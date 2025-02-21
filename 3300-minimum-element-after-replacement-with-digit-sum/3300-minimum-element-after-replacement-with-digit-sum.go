func minElement(nums []int) int {
	ans := 10000
	for _, n := range nums {
		tmp := 0
		for _, i := range strconv.Itoa(n) {
			tmp += int(i - '0')
		}
		if tmp < ans {
			ans = tmp
		}
	}
    
	return ans
}