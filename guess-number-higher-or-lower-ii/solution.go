func umin(current *int, val int) {
    if val < *current {
        *current = val
    }
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func getMoneyAmount(n int) int {
    // T(i, j) 为 [i, j] 的最优解
    // T(i, j) = min{
    // 第一次猜 k, i <= k <= j
    // 猜中，pay k
    // 猜不中，higher ，额外 pay T(i, k-1)
    // 猜不中，lower，额外 pay T(k+1, j)
    // 故三种结果的最坏情况：k + max(T(i, k-1), T(k+1, j))
    // }
    // 边界 T(i, i) = 0，就一个数，不用猜
    
    T := make([][]int, n+2)
    for i := 0; i < n+2; i++ {
        T[i] = make([]int, n+1)
    }
    
    for l := 2; l <= n; l++ {
        for i := 1; i + l - 1 <= n; i++ {
            j := i + l - 1
            v := math.MaxInt32
            for k := i; k <= j; k++ {
                umin(&v, k + max(T[i][k-1], T[k+1][j]))
            }
            T[i][j] = v
        }
    }
    
    return T[1][n]
}
