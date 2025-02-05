func minOperations(nums []int, k int) int {
    counter := 0

    sort.Ints(nums)

    for _, num := range nums {
        if num >= k {
            break
        }

        counter += 1
    }
    
    return counter
}
