func numIdenticalPairs(nums []int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				result++
			}
		}
	}

	return result
}

// func numIdenticalPairs(nums []int) int {
// 	ans := 0
// 	cnt := make(map[int]int)

// 	for _, x := range nums {
// 		ans += cnt[x]
// 		cnt[x]++
//     }

// 	return ans
// }
