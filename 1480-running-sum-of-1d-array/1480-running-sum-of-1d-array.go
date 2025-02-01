func runningSum(nums []int) []int {
    results := make([]int, len(nums))
    temp := 0
    for i, num := range nums {
        results[i] =  num + temp
        temp += num
    }

    return results
}