// func differenceOfSum(nums []int) int {
// 	result := 0

// 	for _, num := range nums {
// 		result += num
// 		for num > 0 {
// 			result -= num % 10
// 			num = num / 10
// 		}
// 	}

// 	return result
// }
func sum(num int) int{
    result := 0
    for num > 0{
        result += num % 10
        num /= 10
    }
    return result
}

func differenceOfSum(nums []int) int {
    result := 0

    for _, num := range nums{
        result += (num - sum(num))
    }
    return result
}