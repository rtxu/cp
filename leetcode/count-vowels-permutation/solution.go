
const MOD = 1e9+7

func countVowelPermutation(n int) int {
    if n <= 0 {
        return 0
    }
    M := map[byte]int{
        'a': 0,
        'e': 1,
        'i': 2,
        'o': 3,
        'u': 4,
    }
    charCount := len(M)
    G := make([][]bool, charCount)
    for i := 0; i < charCount; i++ {
        G[i] = make([]bool, charCount)
    }
    G[0][1] = true
    G[1][0], G[1][2] = true, true
    G[2][0], G[2][1], G[2][3], G[2][4] = true, true, true, true
    G[3][2], G[3][4] = true, true
    G[4][0] = true
    
    // A(i, j) 表示长度为 i，以 j 结尾的序列个数
    // 0 <= i <= n, 0 <= j < 5
    A := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        A[i] = make([]int, charCount)
    }
    for j := 0; j < charCount; j++ {
        A[1][j] = 1
    }
    
    for i := 2; i <= n; i++ {
        for j := 0; j < charCount; j++ {
            for k := 0; k < charCount; k++ {
                if G[k][j] {
                    A[i][j] += A[i-1][k]
                    A[i][j] %= MOD
                }
            }
            
        }
    }
    
    var result int
    for j := 0; j < charCount; j++ {
        result += A[n][j]
        result %= MOD
    }
    return result
}
