func searchInsert(nums []int, target int) int {
    aux := 0
    
    for i, n := range nums {

		if n >= target {
			return i
		}
        
        aux = i
	}
    
	return aux + 1
}
