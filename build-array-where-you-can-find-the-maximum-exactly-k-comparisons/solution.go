const N = 55
const M = 105

const MOD = 1e9+7

func numOfArrays(n int, m int, k int) int {
    /*
    T(i, j, L) 表示序列长度为 i，最大值为 j，最长上升长度为 L 的序列个数
    T(i, j, L) = Sum{
        j * T(i-1, j, L)，第 i 个数 <= j
        T(i-1, sj, L-1)，第 i 个数为最大值 j, 0 <= sj < j 
    }
    则 Sum{T(n, max_value, k)} 为最终答案，其中 1 <= max_value <= m
    */
    var T [N][M][N]int
    T[0][0][0] = 1
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            for L := 1; L <= k; L++ {
                sum := 0
                sum += int(int64(j) * int64(T[i-1][j][L]) % MOD)
                for sj := 0; sj < j; sj++ {
                    sum += T[i-1][sj][L-1]
                    sum %= MOD
                }
                
                T[i][j][L] = sum
            }
        }
    }
    var result int
    for j := 1; j <= m; j++ {
        result += T[n][j][k]
        result %= MOD
    }
    return result
}
