class Solution {
    public String[] findRelativeRanks(int[] score) {
        int n = score.length;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> Integer.compare(b[0], a[0]));

        for (int i = 0; i < n; i++) {
            pq.offer(new int[]{score[i], i});
        }

        int place = 1;
        String[] results = new String[n];
        while (!pq.isEmpty()) {
            int[] top = pq.poll();
            int originalIndex = top[1];
            switch (place) {
                case 1:
                    results[originalIndex] = "Gold Medal";
                    break;

                case 2:
                    results[originalIndex] = "Silver Medal";
                    break;
                case 3:
                    results[originalIndex] = "Bronze Medal";
                    break;
                default:
                    results[originalIndex] = String.valueOf(place);
            }
            place++;
        }

        return results;
    }
}