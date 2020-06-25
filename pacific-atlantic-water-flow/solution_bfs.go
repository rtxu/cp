

var x = []int{0, 1, 0, -1}
var y = []int{1, 0, -1, 0}

type Node struct {
    x, y int
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

func BFS(matrix, flowed [][]int, Q Queue, tag int) {
    m, n := len(matrix), len(matrix[0])
    for !Q.Empty() {
        current := Q.Pop()
        for i := 0; i < 4; i++ {
            var next Node
            next.x, next.y = current.x + x[i], current.y + y[i]
            if next.x >= 0 && next.x < m && next.y >= 0 && next.y < n && 
                matrix[next.x][next.y] >= matrix[current.x][current.y] && 
                flowed[next.x][next.y] & tag == 0 {
                Q.Push(next)
                flowed[next.x][next.y] |= tag
            }
        }
    }
}

func pacificAtlantic(matrix [][]int) [][]int {
    // 1 to pacific, 2 to atlantic, 3 to both
    m := len(matrix)
    if m == 0 {
        return [][]int{}
    }
    n := len(matrix[0])
    
    flowed := make([][]int, m)
    for i := 0; i < m; i++ {
        flowed[i] = make([]int, n)
    }
    
    tag := 1
    Q := make(Queue, 0)
    for i := 0; i < m; i++ {
        if flowed[i][0] & tag == 0 {
            Q.Push(Node{i, 0})
            flowed[i][0] |= tag
        }
    }
    for j := 0; j < n; j++ {
        if flowed[0][j] & tag == 0 {
            Q.Push(Node{0, j})
            flowed[0][j] |= tag
        }
    }
    BFS(matrix, flowed, Q, tag)
    
    tag = 2
    Q = make(Queue, 0)
    for i := 0; i < m; i++ {
        if flowed[i][n-1] & tag == 0 {
            Q.Push(Node{i, n-1})
            flowed[i][n-1] |= tag
        }
    }
    for j := 0; j < n; j++ {
        if flowed[m-1][j] & tag == 0 {
            Q.Push(Node{m-1, j})
            flowed[m-1][j] |= tag
        }
    }
    BFS(matrix, flowed, Q, tag)

    var result [][]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if flowed[i][j] == 3 {
                result = append(result, []int{i, j})
            }
        }
    }
    return result
}
