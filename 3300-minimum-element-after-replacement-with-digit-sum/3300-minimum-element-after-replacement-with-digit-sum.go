func minElement(nums []int) int {
	ans := 0
	for k, n := range nums {
		temp := 0
		for _, i := range strconv.Itoa(n) {
		    ii, _ := strconv.Atoi(string(i))
            temp += ii
		}

        if k == 0 {
			ans = temp
		} else if ans > temp {
			ans = temp
		}
	}

    return ans
}

// 	return ans
// }

// func minElement(nums []int) int {
//     var min int
//     for k, v := range nums {
//         sum := dSum(v)

//         if k == 0 {
//             min = sum
//         } else if min > sum {
//             min = sum
//         }
//     }

//     return min
// }

// func dSum(i int) (sum int) {
//     for _, v := range strconv.Itoa(i) {
//         ii, _ := strconv.Atoi(string(v))

//         sum += ii
//     }

//     return
// }

