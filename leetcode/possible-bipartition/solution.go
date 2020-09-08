func possibleBipartition(N int, dislikes [][]int) bool {
    edges := make([][]int, N+1)
    for i := 0; i < len(dislikes); i++ {
        n1, n2 := dislikes[i][0], dislikes[i][1]
        edges[n1] = append(edges[n1], n2)
        edges[n2] = append(edges[n2], n1)
    }
    colors := make([]int, N+1)
    const NO_COLOR = 0
    const COLOR_1 = 1
    const COLOR_2 = -1
    
    for i := 1; i <= N; i++ {
        if colors[i] == NO_COLOR {
            Q := make([]int, 1)
            Q[0] = i
            colors[i] = COLOR_1
            for len(Q) > 0 {
                current := Q[0]
                Q = Q[1:]
                for e := 0; e < len(edges[current]); e++ {
                    next := edges[current][e]
                    if colors[next] == NO_COLOR {
                        Q = append(Q, next)
                        colors[next] = -colors[current]
                    } else if colors[next] == colors[current] {
                        return false
                    }
                }
            }
        }
    }
    
    return true
}
