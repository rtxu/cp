func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
    G := make([][]bool, n)
    for i := 0; i < n; i++ {
        G[i] = make([]bool, n)
    }
    for i := 0; i < len(prerequisites); i++ {
        from, to := prerequisites[i][0], prerequisites[i][1]
        G[from][to] = true
    }
    
    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if G[i][k] && G[k][j] {
                    G[i][j] = true
                }
            }
        }
    }
    
    queryCnt := len(queries)
    result := make([]bool, queryCnt)
    for i := 0; i < queryCnt; i++ {
        from, to := queries[i][0], queries[i][1]
        result[i] = G[from][to]
    }
    
    return result
}
