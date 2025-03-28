class Solution {
    public int minEatingSpeed(int[] piles, int h) {
        int left = 1, right = 0;
        
        for (int pile : piles) {
            right = Math.max(right, pile);
        }
        
        while (left < right) {
            int mid = left + (right - left) / 2;
            
            if (canAdjust(piles, h, mid)) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        
        return left;
    }

    private static boolean canAdjust(int[] piles, int h, int x) {
        long requiredAdjustments = 0;
        
        for (int pile : piles) {
            requiredAdjustments += (pile + x - 1) / x;
            
            if (requiredAdjustments > h) {
                return false;
            }
        }
        
        return true;
    }
}