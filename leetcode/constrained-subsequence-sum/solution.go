type Node struct {
    val, i int
}

type Queue []Node

func (q Queue) Empty() bool {
    return len(q) == 0
}

func (q *Queue) Push(n Node) {
    *q = append(*q, n)
}

func (q *Queue) PopHead() {
    old := *q
    *q = old[1:]
}

func (q *Queue) PopTail() {
    old := *q
    n := len(old)
    *q = old[:n-1]
}

func (q Queue) Tail() Node {
    return q[len(q)-1]
}

func (q Queue) Head() Node {
    return q[0]
}

func umax(current *int, val int) {
    if val > *current {
        *current = val
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func constrainedSubsetSum(nums []int, k int) int {
    // T(i) 表示截止到 i 位置且 i 位置必选的最优解
    // T(i) = max{max{T(i-1), T(i-k)}, 0} + nums[i]
    n := len(nums)
    T := make([]int, n)
    // monotonic decreasing queue
    q := make(Queue, 0)
    T[0] = nums[0]
    q.Push(Node{T[0], 0})
    for i := 1; i < n; i++ {
        for !q.Empty() && q.Head().i < i-k {
            q.PopHead()
        }
        
        T[i] = max(q.Head().val, 0) + nums[i]
        for !q.Empty() && q.Tail().val <= T[i] {
            q.PopTail()
        }
        q.Push(Node{T[i], i})
    }
    result := -100000009
    for i := 0; i < n; i++ {
        umax(&result, T[i])
    }
    return result
}
