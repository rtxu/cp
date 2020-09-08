// An IntHeap is a min-heap of ints.
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

func nthUglyNumber(n int) int {
    
    h := &IntHeap{1}
    factors := []int{2, 3, 5}
    M := make(map[int]interface{})
    
    v := 0
    for i := 1; i <= n; i++ {
        v = heap.Pop(h).(int)

        for j := 0; j < len(factors); j++ {
            next := v * factors[j]
            if _, exist := M[next]; !exist {
                heap.Push(h, next)
                M[next] = nil
            }
        }
    }
    
    return v
    
}
