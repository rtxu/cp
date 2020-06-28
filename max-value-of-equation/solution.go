// Time: O(n)
// Space: O(n)
// Monotonic Queue

type Node struct {
    Val int
    Pos int
}

type Queue []Node

func (q Queue) Empty() bool { return len(q) == 0 }
func (q Queue) Back() Node { return q[len(q)-1] }
func (q *Queue) PushBack(n Node) {
    *q = append(*q, n)
}
func (q *Queue) PopBack() Node {
    old := *q
    *q = old[:len(old)-1]
    return old[len(old)-1]
}
func (q *Queue) PopFront() Node {
    old := *q
    *q = old[1:]
    return old[0]
}

func findMaxValueOfEquation(points [][]int, k int) int {
    n := len(points)
    end := 0
    Q := make(Queue, 0)
    best := math.MinInt32
    for i := 0; i < n; i++ {
        for end < n && points[end][0] - points[i][0] <= k {
            sum := points[end][0] + points[end][1]
            for !Q.Empty() && Q.Back().Val < sum {
                Q.PopBack()
            }
            Q.PushBack(Node{sum, end})
            end++ 
        }
        for !Q.Empty() && Q[0].Pos <= i {
            Q.PopFront()
        }
        if end-i == 1 {
            // only self in the valid window
        } else {
            v := -points[i][0] + points[i][1] + Q[0].Val
            if v > best {
                best = v
            }
        }
    }
    return best
}
