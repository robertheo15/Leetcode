func mySqrt(x int) int {
	start := 1
	end := x
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	
	return mid
}