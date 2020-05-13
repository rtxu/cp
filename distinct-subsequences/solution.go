func numDistinct(s string, t string) int {
    // A(i, j) 表示 s[0:i] t[0:j] 的答案，则 A(len(s), len(t)) 为最终答案
    // A(i, 0) = 1, 0 <= i < len(s), 空序列算 1 个
    // A(0, j) = 0, 0 < j < len(t)
    // A(i, j) = A(i-1, j) 不以 s[i-1] 结尾的子序列个数
    //          + A(i-1, j-1) 当 s[i-1] == t[j-1] 时，以 s[i-1] 结尾的子序列个数
    ns, nt := len(s), len(t)
    A := make([][]int, ns+1)
    for i := 0; i <= ns; i++ {
        A[i] = make([]int, nt+1)
        A[i][0] = 1
    }
    
    for i := 1; i <= ns; i++ {
        for j := 1; j <= nt; j++ {
            A[i][j] = A[i-1][j]
            if s[i-1] == t[j-1] {
                A[i][j] += A[i-1][j-1]
            }
        }
    }
    return A[ns][nt]
}
