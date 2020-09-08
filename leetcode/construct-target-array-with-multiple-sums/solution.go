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
    
    /*
    sort.Ints(target)
    for i := 0; i < n; i++ {
        if target[i] < 1 {
            return false
        }
        if i > 0 && target[i] > 1 && target[i] == target[i-1] {
            return false
        }
    }
    */
    // 思路：找到最大值，将其减去其他值，获得该值上一轮的取值
    // 考虑 [1, 10^9] 时，做了大量无用的重复计算，为了避免此种情况
    // 维护 max, secondMax, noMaxSum
    // 每次确保 max 值减去 x 个 noMaxSum 后 < secondMax
    // 即 max - x*noMaxSum < secondMax
    // => x > (max - secondMax) / noMaxSum
    
    target2 := IntHeap(target)
    heap.Init(&target2)
    max := heap.Pop(&target2).(int)
    noMaxSum := int64(0)
    for i := 0; i < len(target2); i++ {
        noMaxSum += int64(target2[i])
    }
    secondMax := target2[0]
    
    for !(max == 1 && secondMax == 1) {
        x := int64(max - secondMax) / noMaxSum
        if secondMax > 1 && int64(max - secondMax) % noMaxSum == 0 {
            // secondMax 不为 1 的情况下，减法运算后，max = secondMax
            return false
        }
        if int64(max - secondMax) % noMaxSum > 0 {
            x++
        }
        
        v := max - int(x * noMaxSum)
        if v < 1 {
            return false
        }
        heap.Push(&target2, v)
        max = heap.Pop(&target2).(int)
        secondMax = target2[0]
        noMaxSum = noMaxSum - int64(max) + int64(v)
        
        //fmt.Println("max", max, "secondMax", secondMax, "noMaxSum", noMaxSum, target2)
    }
    
    return true
}
