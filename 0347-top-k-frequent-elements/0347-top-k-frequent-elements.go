// Element represents an element in the priority queue
type Element struct {
	Value     int
	Frequency int
}

// MinHeap is a min-heap implementation for Element
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: Count the frequency of each element using a hashmap
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Step 2: Create a min-heap (priority queue) to keep track of the k most frequent elements
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Step 3: Traverse the frequency map and add elements to the min-heap
	for num, frequency := range frequencyMap {
		heap.Push(minHeap, Element{Value: num, Frequency: frequency})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	// Step 4: Retrieve the result from the min-heap
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(minHeap).(Element).Value
	}

	return result
}