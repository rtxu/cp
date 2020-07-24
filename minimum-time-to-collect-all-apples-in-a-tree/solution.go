func postOrder(current, parent int, G [][]int, hasApple []bool, result *int) int {
    appleCnt := 0
    if hasApple[current] {
        appleCnt++
    }
    for i := 0; i < len(G[current]); i++ {
        to := G[current][i]
        if to != parent {
            childAppleCnt := postOrder(to, current, G, hasApple, result)
            if childAppleCnt > 0 {
                *result += 2
            }
            appleCnt += childAppleCnt
        }
    }
    return appleCnt
}

func minTime(n int, edges [][]int, hasApple []bool) int {
    G := make([][]int, n)
    for i := 0; i < n; i++ {
        G[i] = make([]int, 0)
    }
    m := len(edges)
    for i := 0; i < m; i++ {
        a, b := edges[i][0], edges[i][1]
        G[a] = append(G[a], b)
        G[b] = append(G[b], a)
    }
    
    var result int
    postOrder(0, -1, G, hasApple, &result)
    return result
}

