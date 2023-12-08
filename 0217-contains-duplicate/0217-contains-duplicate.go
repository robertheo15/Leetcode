func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sort.Ints(nums)
    set := make(map[int]struct{})
	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}