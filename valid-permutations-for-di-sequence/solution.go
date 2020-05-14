
const MOD = 1e9 + 7

// ref: https://leetcode.com/problems/valid-permutations-for-di-sequence/discuss/196939/Easy-to-understand-solution-with-detailed-explanation
func numPermsDISequence(S string) int {
    // A(i, j) 表示满足 S[0:i+1] 数字为 0...i+1 的排列数，且该这些排列以 j 结尾
    // 0 <= i < len(S), 0 <= j <= i+1
    // 最终结果 = Sum(A(len(S)-1, k)), 0 <= k <= len(S)
    // 则 A(i, j) 为
    // 1) 当 S[i] = 'D' 时，Sum(A(i-1, k)), j <= k < i
    // 2) 当 S[i] = 'I' 时，Sum(A(i-1, k)), 0 <= k < j
    
    n := len(S)
    if n == 0 {
        return 0
    }
    A := make([][]int, n)
    for i := 0; i < n; i++ {
        A[i] = make([]int, i+2)
    }
    
    if S[0] == 'D' {
        A[0][0] = 1
    } else {
        A[0][1] = 1
    }
    
    for i := 1; i < n; i++ {
        if S[i] == 'D' {
            for j := 0; j <= i+1; j++ {
                sum := 0
                for k := j; k <= i; k++ {
                    sum = (sum + A[i-1][k]) % MOD
                }
                A[i][j] = sum
            }
        } else {
            for j := 0; j <= i+1; j++ {
                sum := 0
                for k := 0; k < j; k++ {
                    sum = (sum + A[i-1][k]) % MOD
                }
                A[i][j] = sum
            }
        }
    }
    
    result := 0
    for i := 0; i <= n; i++ {
        result += A[n-1][i]
    }
    result %= MOD
    return result
}
