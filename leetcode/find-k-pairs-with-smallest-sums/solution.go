
type sum struct {
    i, j, val int
}

type MinValHeap []sum

func (a MinValHeap) Len() int            { return len(a) }
func (a MinValHeap) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a MinValHeap) Less(i, j int) bool  { return a[i].val < a[j].val }

func (a *MinValHeap) Push(v interface{}) {
    *a = append(*a, v.(sum))
}

func (a *MinValHeap) Pop() interface{} {
    old := *a
    n := len(old)
    v := old[n-1]
    *a = old[0:n-1]
    return v
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    isSwap := false
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
        isSwap = true
    }
    n1, n2 := len(nums1), len(nums2)
    if n1 == 0 || n2 == 0 {
        return [][]int{}
    }
    heapSize := min(n1, k)
    sumHeap := make(MinValHeap, 0, heapSize)
    // fmt.Println("n1", n1, "n2", n2, "k", k)
    for i := 0; i < heapSize; i++ {
        heap.Push(&sumHeap, sum{i, 0, nums1[i]+nums2[0]})
    }
    var result [][]int
    for k > 0 && sumHeap.Len() > 0 {
        current := heap.Pop(&sumHeap).(sum)
        if isSwap {
            result = append(result, []int{nums2[current.j], nums1[current.i]})
        } else {
            result = append(result, []int{nums1[current.i], nums2[current.j]})
        }
        if current.j++; current.j < n2 {
            heap.Push(&sumHeap, sum{current.i, current.j, nums1[current.i]+nums2[current.j]})
        }
        k--
    }
    return result
}
