class Solution {
    public int smallestDivisor(int[] nums, int threshold) {
        int left = 1, right = 0;
        
        for (int num : nums) {
            right = Math.max(right, num);
        }
        
        while (left < right) {
            int mid = left + (right - left) / 2;
            
            if (canAdjust(nums, threshold, mid)) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        
        return left;
    }

    private static boolean canAdjust(int[] nums, int threshold, int x) {
        long requiredAdjustments = 0;
        
        for (int num : nums) {
            requiredAdjustments += (num + x - 1) / x;
            
            if (requiredAdjustments > threshold) {
                return false;
            }
        }
        
        return true;
    }
}