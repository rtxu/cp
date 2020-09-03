type Node struct {
    i, j int
}

type Queue []Node

func (q Queue) Empty() bool {
    return len(q) == 0
}

func (q *Queue) Push(n Node) {
    *q = append(*q, n)
}

func (q *Queue) Pop() Node {
    old := *q
    *q = old[1:]
    return old[0]
}

func findDiagonalOrder(nums [][]int) []int {
    visited := make([][]bool, len(nums))
    for i := 0; i < len(nums); i++ {
        visited[i] = make([]bool, len(nums[i]))
    }
    q := make(Queue, 0)
    q.Push(Node{0, 0})
    visited[0][0] = true
    var result []int
    for !q.Empty() {
        current := q.Pop()
        result = append(result, nums[current.i][current.j])
        ni, nj := current.i+1, current.j
        if ni < len(nums) && nj < len(nums[ni]) && !visited[ni][nj] {
            q.Push(Node{ni, nj})
            visited[ni][nj] = true
        }
        ni, nj = current.i, current.j+1
        if nj < len(nums[ni]) && !visited[ni][nj] {
            q.Push(Node{ni, nj})
            visited[ni][nj] = true
        }
        
    }
    return result
}
