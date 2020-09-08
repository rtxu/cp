const MOD = 1e9+7

func countPalindromicSubsequences(S string) int {
    n := len(S)
    // A(i, j), i < j 为 S[i:j] 的答案，则 A(0, n) 为最终答案
    // A(i, j) = S[i:j] 中包含的不同字符数 + Na_a + Nb_b + Nc_c + Nd_d
    // 其中 Na_a 表示 a_a 形式的回文串数量，则 Na_a = A(ia+1, ja) + 1，1 为空串
    A := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        A[i] = make([]int, n+1)
    }
    
    for i := n-1; i >= 0; i-- {
        uniqCharCount := 0
        M := make(map[byte]bool)
        begin, end := make(map[byte]int), make(map[byte]int)
        
        for j := i+1; j <= n; j++ {
            currentChar := S[j-1]
            if _, exist := M[currentChar]; !exist {
                uniqCharCount++
                M[currentChar] = true
                begin[currentChar] = j-1
            }
            end[currentChar] = j-1
            A[i][j] += uniqCharCount
            //fmt.Printf("A[%d][%d] = %d ", i, j, A[i][j])
            for char, charBegin := range begin {
                charEnd := end[char]
                if charBegin < charEnd {
                    count := 1 // 1 for empty sequence
                    count += A[charBegin+1][charEnd]
                    A[i][j] += count
                }
            }
            A[i][j] %= MOD
            //fmt.Printf("A[%d][%d] = %d \n", i, j, A[i][j])
        }
    }
    
    return A[0][n]
}
