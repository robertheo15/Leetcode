func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return false
	}

	_, duplicate := sliceToMap(nums)

	return duplicate
}

func sliceToMap(nums []int) (map[int]int, bool) {
	numMap := make(map[int]int)

	for _, num := range nums {

		if _, exist := numMap[num]; !exist {
			numMap[num] = 0
		}

		numMap[num] += 1

		if numMap[num] > 1 {
			return numMap, true
		}
	}

	return numMap, false
}