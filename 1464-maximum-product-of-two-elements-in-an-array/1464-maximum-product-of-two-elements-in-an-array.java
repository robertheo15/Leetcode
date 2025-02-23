class Solution {
    public int maxProduct(int[] nums) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(new Comparator<Integer>() {
            public int compare(Integer a, Integer b) {
                return b - a;
            }
        });

        for (int num : nums) {
            pq.offer(num);
        }

        int i = pq.poll();
        int j = pq.poll();

        return (i-1) * (j-1);
    }
}