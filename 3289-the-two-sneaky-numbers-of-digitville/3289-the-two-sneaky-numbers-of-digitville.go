func getSneakyNumbers(nums []int) []int {
    res := make([]int, 0)
    counts := make([]int, 101)

    for _, num := range nums {
        counts[num]++
        if counts[num] > 1 {
            res = append(res, num)
        }
    }

    return res
}