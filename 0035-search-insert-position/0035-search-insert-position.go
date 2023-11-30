func searchInsert(nums []int, target int) int {
    for index, val := range nums {
        if val == target || val > target  {
            return index
        }
    }

    return len(nums)
}