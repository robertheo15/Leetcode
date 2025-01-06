func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	secondLast, last := 0, 1
	for i := 1; i <= n; i++ {
		secondLast, last = last, secondLast+last
	}
	return last
}