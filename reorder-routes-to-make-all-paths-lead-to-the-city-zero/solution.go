func minReorder(n int, connections [][]int) int {
    edge := make([][]int, n)
    for j := 0; j < len(connections); j++ {
        from, to := connections[j][0], connections[j][1]
        edge[from] = append(edge[from], -to)
        edge[to] = append(edge[to], from)
    }
    
    Q := make([]int, n)
    visited := make([]bool, n)
    Q[0] = 0
    visited[0] = true
    qHead, qTail := 0, 1
    
    var result int
    for qTail > qHead {
        current := Q[qHead]
        qHead++
        for e := 0; e < len(edge[current]); e++ {
            j := edge[current][e]
            var to int
            if j < 0 {
                to = -j
            } else {
                to = j
            }
            if !visited[to] {
                Q[qTail] = to
                qTail++
                visited[to] = true
                if j < 0 {
                    //fmt.Println(current, j)
                    result++
                }
            }
        }
    }
    
    return result
}

/*
5
[[1,0],[1,2],[3,2],[3,4]]
3
[[1,0],[2,0]]
*/
