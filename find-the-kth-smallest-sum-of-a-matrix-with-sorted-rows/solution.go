
type Node struct {
    sum int
    state []int
}

func (n Node) Key() string {
    var sb strings.Builder
    for i := 0; i < len(n.state); i++ {
        sb.WriteString(fmt.Sprintf("%d-", n.state[i]))
    }
    return sb.String()
}

// An IntHeap is a min-heap of ints.
type Heap []Node

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(mat [][]int, k int) int {
    m, n := len(mat), len(mat[0])
    h := make(Heap, 0)
    initial := Node{}
    for i := 0; i < m; i++ {
        initial.sum += mat[i][0]
        initial.state = append(initial.state, 0)
    }
    heap.Push(&h, initial)
    visited := make(map[string]bool)
    visited[initial.Key()] = true
    for i := 2; i <= k; i++ {
        current := heap.Pop(&h).(Node)
        // fmt.Println("current", current.sum, current.state)
        for j := 0; j < m; j++ {
            if current.state[j] + 1 < n {
                next := Node{
                    sum : current.sum,
                    state : make([]int, m),
                }
                copy(next.state, current.state)
                next.sum -= mat[j][next.state[j]]
                next.state[j]++
                next.sum += mat[j][next.state[j]]
                if !visited[next.Key()] {
                    heap.Push(&h, next)
                    visited[next.Key()] = true
                }
                
                // fmt.Println("next", next.sum, next.state)
            }
        }
    }
    return heap.Pop(&h).(Node).sum
}
