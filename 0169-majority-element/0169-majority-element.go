func majorityElement(nums []int) int {
	numsMap := sliceToMap(nums)
	max := 0
	res := 0
	for key, num := range numsMap {
		if num > max {
			max = num
			res = key
		}
	}

	return res
}

func sliceToMap(nums []int) map[int]int {
	numsMap := make(map[int]int)

	for _, num := range nums {
		if _, exist := numsMap[num]; !exist {
			numsMap[num] = 1
			continue
		}

		numsMap[num] += 1
	}

	return numsMap
}
