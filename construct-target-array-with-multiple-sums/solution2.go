// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func isPossible(target []int) bool {
    n := len(target)
    if n == 1 {
        return target[0] == 1
    }
    
    // 思路：找到最大值，将其减去其他值之和，获得该值上一轮的取值
    // 考虑 [1, 10^9] 时，做了大量无用的重复计算，为了避免此种情况，max %= noMaxSum
    
    target2 := IntHeap(target)
    heap.Init(&target2)
    sum := 0
    for i := 0; i < len(target2); i++ {
        sum += target2[i]
    }
    
    for {
        max := heap.Pop(&target2).(int)
        sum -= max
        if max == 1 || sum == 1 {
            return true
        }
        if max < sum || sum == 0 || max % sum == 0 {
            return false
        }
        max %= sum
        heap.Push(&target2, max)
        sum += max
    }
    
    return true
}
