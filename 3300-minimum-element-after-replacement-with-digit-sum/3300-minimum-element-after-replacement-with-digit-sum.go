// func minElement(nums []int) int {
// 	ans := 10000
// 	for _, n := range nums {
// 		tmp := 0
// 		for _, i := range strconv.Itoa(n) {
// 			tmp += int(i - '0')
// 		}
// 		if tmp < ans {
// 			ans = tmp
// 		}
// 	}

// 	return ans
// }
func minElement(nums []int) int {
    var min int
    for k, v := range nums {
        sum := dSum(v)

        if k == 0 {
            min = sum
        } else if min > sum {
            min = sum
        }
    }

    return min
}

func dSum(i int) (sum int) {
    for _, v := range strconv.Itoa(i) {
        ii, _ := strconv.Atoi(string(v))

        sum += ii
    }

    return
}