func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func minimumDeleteSum(s1 string, s2 string) int {
    // A(i, j) s1[0:i] s2[0:j] 的最小值，则 A(len(s1), len(s2)) 为最终结果
    // 初始条件：
    // 1) A(0, 0) = 0
    // 2) A(0, j) = Sum(s2[0]...s2[j-1]), 1 <= j <= len(s2)
    // 3) A(i, 0) = Sum(s1[0]...s1[i-1]), 1 <= i <= len(s1)
    // 转移公式：
    // A(i, j) = min{
    // 1) 当 s1[i-1] == s2[j-1] 时无需删除任何字符, A(i-1, j-1)
    // 2) 删掉 s1 子串的最后一个字符，A(i-1, j) + s1[i-1]
    // 3) 删掉 s2 子串的最后一个字符，A(i, j-1) + s2[j-1]
    // }
    
    n1, n2 := len(s1), len(s2)
    A := make([][]int, n1+1)
    for i := 0; i <= n1; i++ {
        A[i] = make([]int, n2+1)
    }
    for i := 1; i <= n1; i++ {
        A[i][0] = A[i-1][0] + int(s1[i-1])
    }
    for j := 1; j <= n2; j++ {
        A[0][j] = A[0][j-1] + int(s2[j-1])
    }
    
    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            A[i][j] = min(A[i-1][j] + int(s1[i-1]), A[i][j-1] + int(s2[j-1]))
            if s1[i-1] == s2[j-1] {
                A[i][j] = min(A[i][j], A[i-1][j-1])
            }
        }
    }
    
    return A[n1][n2]
}
