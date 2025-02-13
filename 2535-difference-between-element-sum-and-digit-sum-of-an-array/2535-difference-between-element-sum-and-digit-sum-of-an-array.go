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
    if num < 10 {
        return num
    }

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
// func differenceOfSum(nums []int) int {
// 	// DONE-- loop iterate to nums[]
// 	// DONE-- take the sum of all digits in array
// 	// DONE-- take the sum of all digits
// 	// DONE-- write down a function for this sum operation
// 	// DONE-- return the result
// 	Arraysum := 0
// 	Digitsum := 0
// 	for _, v := range nums {
// 		Arraysum += v
// 		Digitsum += digitSum(v)
// 	}
// 	return Arraysum - Digitsum

// }

// func digitSum(n int) int {
// 	if n < 10 {
// 		return n
// 	}
// 	var sum int

// 	for n > 0 {
// 		sum += n % 10
// 		n = n / 10
// 	}
// 	return sum
// }