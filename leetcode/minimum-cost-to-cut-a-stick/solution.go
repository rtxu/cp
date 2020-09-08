
func minCost(n int, cuts []int) int {
    // T(i, j) 表示 [i, j] 的最优解
    // T(0, n) 为最终解
    // 注意：第一刀位置确定，则 cuts 的分配确定
    // T(0, n) = min(n + T(0, k) + T(k, n)) for all k, k 属于 cuts
    
    M := make(map[int]map[int]int)
    return solve(0, n, cuts, M)
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func solve(i, j int, cuts []int, M map[int]map[int]int) int {
    if _, exist := M[i]; exist {
        if _, exist := M[i][j]; exist {
            return M[i][j]
        }
    } else {
        M[i] = make(map[int]int)
    }
    
    cost := math.MaxInt32
    n := len(cuts)
    for k := 0; k < n; k++ {
        if cuts[k] > i && cuts[k] < j {
            cost = min(cost, j-i + solve(i, cuts[k], cuts, M) + solve(cuts[k], j, cuts, M))
        }
    }
    if cost == math.MaxInt32 {
        cost = 0
    }
    
    M[i][j] = cost
    // fmt.Println(i, j, cost)
    return cost;
}
