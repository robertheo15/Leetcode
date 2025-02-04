func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    
    cnt := 0
    for i := 0; i < len(nums1); i++ {
        n := nums1[i]
        for j := 0; j < len(nums2); j++ {
            m := nums2[j] * k
            if n % m == 0 {
                cnt ++
            }
        }
    }

    return cnt
}