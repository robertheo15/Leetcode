func getConcatenation(nums []int) []int {
    for i,_ := range nums{
        nums = append(nums, nums[i])
    }
    return nums
}