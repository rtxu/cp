// Time: O(len(s1)*len(s2))
// Space: O(len(s1)*len(s2))

func backtrack(s1, s2, s3 string, i1, i2, i3 int) bool {
    if i3 == len(s3) {
        return true
    }
    if i1 < len(s1) && s3[i3] == s1[i1] {
        if backtrack(s1, s2, s3, i1+1, i2, i3+1) {
            return true
        }
    }
    if i2 < len(s2) && s3[i3] == s2[i2] {
        if backtrack(s1, s2, s3, i1, i2+1, i3+1) {
            return true
        }
    }
    return false
}


func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }
    
    // A(i, j) 表示 s1[0:i] 和 s2[0:j] 是否可以形成 s3[0:i+j]
    // 初始条件：A(0, 0) = true
    // 转移方程：A(i, j) = (A(i-1, j) && s1[i-1] == s3[i+j-1]) || (A(i, j-1) && s2[j-1] == s3[i+j-1])
    A := make([][]bool, len(s1)+1)
    for i := 0; i < len(s1)+1; i++ {
        A[i] = make([]bool, len(s2)+1)
    }
    A[0][0] = true
    for i := 1; i <= len(s1); i++ {
        A[i][0] = A[i-1][0] && s1[i-1] == s3[i-1]
    }
    for j := 1; j <= len(s2); j++ {
        A[0][j] = A[0][j-1] && s2[j-1] == s3[j-1]
    }
    for i := 1; i <= len(s1); i++ {
        for j := 1; j <= len(s2); j++ {
            A[i][j] = (A[i-1][j] && s1[i-1] == s3[i+j-1]) || (A[i][j-1] && s2[j-1] == s3[i+j-1])
        }
    }
    
    return A[len(s1)][len(s2)]
}
