const MOD = 1e9+7

func profitableSchemes(G int, P int, group []int, profit []int) int {
    // A(i, j) 以 i 为人员成本，以 j 为利润的方案数
    // 0 <= i <= G, 0 <= j <= P，对于所有 >P 的利润，记为 P
    // 初始条件：A(0, 0) = 1
    A := make([][]int, G+1)
    for i := 0; i <= G; i++ {
        A[i] = make([]int, P+1)
    }
    A[0][0] = 1
    
    var result int
    n := len(group)
    for i := 0; i < n; i++ {
        for g := G; g >= 0; g-- {
            for p := P; p >= 0; p-- {
                if A[g][p] > 0 {
                    next_g := g+group[i]
                    next_p := p+profit[i]
                    if next_g > G {
                        continue
                    }
                    if next_p > P {
                        next_p = P
                    }
                    A[next_g][next_p] += A[g][p]
                    A[next_g][next_p] %= MOD
                    if next_p >= P {
                        result += A[g][p]
                        result %= MOD
                    }
                }
            }
        }
    }
    
    return result
}
