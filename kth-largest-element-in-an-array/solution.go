
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
    n := len(nums)
    h := make(IntHeap, 0)
    for i := 0; i < k; i++ {
        heap.Push(&h, nums[i])
    }
    for i := k; i < n; i++ {
        if nums[i] > h[0] {
            heap.Pop(&h)
            heap.Push(&h, nums[i])
        }
    }
    return h[0]
}
