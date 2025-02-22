class Solution {
    public long maxKelements(int[] nums, int k) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(new Comparator<Integer>() {
            public int compare(Integer a, Integer b) {
                return  b - a;
            }
        });

        for (int num: nums) {
            pq.offer(num);
        }

        long result = 0;

        for(int i = 0; i < k ; i++) {
            int value = pq.poll();
            result += value;
            pq.offer((int) Math.ceil(value / 3.0));
        }

        return result;
    }
}