const MOD = 1e9+7

func numberWays(hats [][]int) int {
    // T(i, s) 表示前 i 个帽子戴给集合 s 的方案数
    // 则 T(40, 1<<n - 1) 为最终答案
    // T(0, empty_set) = 1
    // T(i, s) = T(i-1, s) + Sum{T(i-1, s-j) 将第 i 顶帽子给 j (即 i in hats[j])}
    n := len(hats)
    T := make([][]int, 41)
    for i := 0; i < 41; i++ {
        T[i] = make([]int, 1<<n)
    }
    T[0][0] = 1
    for i := 1; i < 41; i++ {
        for s := 0; s < 1<<n; s++ {
            cnt := T[i-1][s]
            for j := 0; j < n; j++ {
                if s & (1 << j) > 0 {
                    for k := 0; k < len(hats[j]); k++ {
                        if i == hats[j][k] {
                            cnt += T[i-1][s ^ (1 << j)]
                            cnt %= MOD
                        }
                    }
                }
            }
            T[i][s] = cnt
        }
    }
    return T[40][1<<n-1]
}
