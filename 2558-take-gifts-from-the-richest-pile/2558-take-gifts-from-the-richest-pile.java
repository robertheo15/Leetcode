class Solution {
    public long pickGifts(int[] gifts, int k) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(new Comparator<Integer>() {
            public int compare(Integer a, Integer b) {
                return b - a;
            }
        });

        for (int gift : gifts) {
            pq.offer(gift);
        }

        for (int i = 0; i < k; i++) {
            int gift = pq.poll();
            pq.offer((int)Math.floor(Math.sqrt(gift)));
        }

        long ans = 0;
        while (!pq.isEmpty()) {
            ans += pq.poll();
        }

        return ans;
    }
}