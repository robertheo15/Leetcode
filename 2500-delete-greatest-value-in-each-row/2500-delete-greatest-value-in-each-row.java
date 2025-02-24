class Solution {
    public int deleteGreatestValue(int[][] grid) {
    	ArrayList<PriorityQueue<Integer>> list=new ArrayList<>();
    	int m = grid.length, n = grid[0].length;

    	for (int i = 0; i < grid.length ; i++) {
    		PriorityQueue<Integer> pq=new PriorityQueue<>(Collections.reverseOrder());
            for (int j = 0; j < grid[0].length ; j++) {
                pq.add(grid[i][j]);
            }

            list.add(pq);
        }

        int result = 0;
    	for (int i = 0; i < n ; i++) {
    		int num=-1;
    		for (int j = 0 ; j < m ; j++) {
    			PriorityQueue<Integer> pq=list.get(j);
    			int temp = pq.poll();
					
    			num = Math.max(num, temp);
    		}
    		result+=num;
    	}

    	return result;
    }
}