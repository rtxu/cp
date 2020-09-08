func partition(s string) [][]string {
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
    
    // result[i] 表示遍历到第 i-1 个字符时的结果集
    result := make([][][]string, n+1)
    result[0] = make([][]string, 1)
    result[0][0] = []string{}
    for i := 1; i <= n; i++ {
        for j := i; j >= 1; j-- {
            if P[j-1][i-1] {
                for k := 0; k < len(result[j-1]); k++ {
                    oldResult := result[j-1][k]
                    newResult := make([]string, len(oldResult) + 1)
                    copy(newResult, oldResult)
                    newResult[len(oldResult)] = s[j-1:i]
                    result[i] = append(result[i], newResult)
                }
            }
        }
    }
    
    return result[n]
}
