class Solution {
    public int minimumOperations(int[] nums) {
        int ans = 0;
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>((a, b) -> b - a);
        while (true) {
            for (int i : nums) {
                if (i != 0)
                    maxHeap.offer(i);
            }

            if (maxHeap.isEmpty()) {
                break;
            }

            int minEle = maxHeap.poll();
            for (int i = 0; i < nums.length; i++) {
                if (nums[i] != 0) {
                    nums[i] = nums[i] - minEle;
                }
            }

            ans++;

            if (!maxHeap.isEmpty()) {
                maxHeap.clear();
            }
        }

        return ans;
    }
}