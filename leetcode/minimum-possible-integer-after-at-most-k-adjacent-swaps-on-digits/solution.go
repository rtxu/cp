
type BIT []int

func (bit BIT) update(i int, delta int) {
    for ; i < len(bit); i += i&-i {
        bit[i] += delta
    }
}

func (bit BIT) prefixSum(i int) int {
    var result int
    for ; i > 0; i -= i&-i {
        result += bit[i]
    }
    return result
}

type Queue []int

func (q Queue) Empty() bool {
    return len(q) == 0
}
    
func (q *Queue) Push(v int) {
    *q = append(*q, v)
}

func (q *Queue) Pop() int {
    old := *q
    *q = old[1:]
    return old[0]
}

func (q Queue) Top() int {
    return q[0]
}

func minInteger(num string, k int) string {
    var queue [10]Queue
    for i := 0; i < 10; i++ {
        queue[i] = make(Queue, 0)
    }
    n := len(num)
    for i := 0; i < n; i++ {
        digit := num[i] - '0'
        queue[digit].Push(i+1)
    }
    
    selected := make([]bool, n+1)
    bit := make(BIT, n+1)
    for i := 1; i <= n; i++ {
        bit.update(i, 1)
    }
    
    var result strings.Builder
    for j := 1; j <= n && k > 0; j++ {
        for i := 0; i < 10 && k > 0; i++ {
            if !queue[i].Empty() {
                current := queue[i].Top()
                moveCnt := bit.prefixSum(current) - 1
                if moveCnt <= k {
                    selected[current] = true
                    result.WriteByte(byte('0' + i))
                    // fmt.Println("i", i, "current", current, "moveCnt", moveCnt, "k", k, result.String())
                    k -= moveCnt
                    bit.update(current, -1)
                    queue[i].Pop()
                    
                    break
                }
            }
        }
    }
    
    for i := 1; i <= n; i++ {
        if !selected[i] {
            result.WriteByte(num[i-1])
        }
    }
    
    return result.String()
}

