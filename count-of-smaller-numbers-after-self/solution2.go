type Node struct {
    Val int
    Pos int
    Count int
}

func countSmaller(nums []int) []int {
    n := len(nums)
    nodes := make([]Node, n)
    for i := 0; i < n; i++ {
        nodes[i] = Node{nums[i], i, 0}
    }
    
    mergeSort(nodes)
    counts := make([]int, n)
    for i := 0; i < n; i++ {
        counts[nodes[i].Pos] = nodes[i].Count
    }
    return counts
}

func mergeSort(nums []Node) {
    n := len(nums)
    if n <= 1 {
        return
    }
    
    mid := n >> 1
    mergeSort(nums[:mid])
    mergeSort(nums[mid:])
    
    for i, j := 0, mid; i < mid; i++ {
        for j < n && nums[j].Val < nums[i].Val { j++ }
        nums[i].Count += j - mid
    }
    
    inplaceMerge(nums, mid)
}

func inplaceMerge(nums []Node, mid int) {
    n := len(nums)
    buf := make([]Node, n)
    bufLen := 0
    for i, j := 0, mid; i < mid; i++ {
        for j < n && nums[j].Val < nums[i].Val {
            buf[bufLen] = nums[j]
            bufLen++
            j++
        }
        buf[bufLen] = nums[i]
        bufLen++
    }
    copy(nums, buf[:bufLen])
}
