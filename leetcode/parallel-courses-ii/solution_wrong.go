type Node struct {
    id          int
    outdegree   int
}

type MaxHeap []Node

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].outdegree > h[j].outdegree }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minNumberOfSemesters(n int, deps [][]int, k int) int {
    indegree := make([]int, n+1)
    outdegree := make([]int, n+1)
    depsCnt := len(deps)
    for i := 0; i < depsCnt; i++ {
        x, y := deps[i][0], deps[i][1]
        indegree[y]++
        outdegree[x]++
    }

    
    Q := make(MaxHeap, 0)
    for i := 1; i <= n; i++ {
        if indegree[i] == 0 {
            heap.Push(&Q, Node{i, outdegree[i]})
        }
    }
    count := 0
    for Q.Len() > 0 {
        qsize := Q.Len()
        count += qsize / k
        if qsize < k {
            count++
        } else {
            qsize -= qsize % k
        }
        nextQ := []int{}
        for i := 0; i < qsize; i++ {
            current := heap.Pop(&Q).(Node)
            fmt.Println("current", current)
            for j := 0; j < depsCnt; j++ {
                x, y := deps[j][0], deps[j][1]
                if x == current.id {
                    indegree[y]--
                    if indegree[y] == 0 {
                        nextQ = append(nextQ, y)
                    }
                }
            }
        }
        for i := 0; i < len(nextQ); i++ {
            y := nextQ[i]
            heap.Push(&Q, Node{y, outdegree[y]})
        }
    }
    return count
}
