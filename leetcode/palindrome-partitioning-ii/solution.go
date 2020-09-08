func minCut(s string) int {
    // P(i, j) 表示 s[i:j+1] 是否为回文串, 0 <= i < n, i <= j < n
    // 显然
    // 1) P(i，i) 皆为回文
    // 2) 空串，j < i 也为回文
    
    n := len(s)
    P := make([][]bool, n)
    for i := 0; i < n; i++ {
        P[i] = make([]bool, n)
    }
    
    for i := n-1; i >= 0; i-- {
        P[i][i] = true
        for j := i+1; j < n; j++ {
            if s[i] == s[j] && (j-1 < i+1 || P[i+1][j-1]) {
                P[i][j] = true
            }
        }
        //fmt.Println(P[i])
    }
    
    // C[i] 表示截止到 i 位置（不包含）的最小 cut 数
    // 则 C[n] 为最终答案
    // C[i] = min(C[k]+1), k 为任意可以与 i-1 构成回文的位置
    C := make([]int, n+1)
    C[0], C[1] = -1, 0
    for j := 2; j <= n; j++ {
        pEnd := j-1
        v := -1
        for pStart := pEnd; pStart >= 0; pStart-- {
            if P[pStart][pEnd] {
                if v == -1 || C[pStart]+1 < v {
                    v = C[pStart]+1
                }
            }
        }
        C[j] = v
    }
    
    return C[n]
}

