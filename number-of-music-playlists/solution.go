const MOD = 1e9+7

func numMusicPlaylists(N int, L int, K int) int {
    // A(i, j) 长度为 i 的序列，使用了 j 种歌曲的方案数，1 <= i <= L, 1 <= j <= N
    // 则 A(L, N) 为最终答案
    // A(i, 1) = N, 1 <= i <= L
    // A(1, j) = 0, j > 1
    // A(i, j) = 
    // 1) 如果第 j 首歌为新歌，A(i-1, j-1) * (N-(j-1))
    // 2) 如果第 j 首歌为老歌（不考虑 k 的限制），A(i-1, j) * j
    // 考虑 k 的限制
    // 1) 不变
    // 2) 因为 k 的存在，i-1 歌单的最后 k 首歌必不相同，故第 j 首歌可用歌曲种类为 j-k，当 j <= k 时，无歌可用
    A := make([][]int, L+1)
    for i := 1; i <= L; i++ {
        A[i] = make([]int, N+1)
    }
    A[1][1] = N
    for i := 2; i <= L; i++ {
        for j := 1; j <= N; j++ {
            v := int64(0)
            v += int64(A[i-1][j-1]) * int64(N-(j-1))
            if j > K {
                v += int64(A[i-1][j]) * int64(j-K)
            }
            v %= MOD
            A[i][j] = int(v)
        }
        //fmt.Println(A[i])
    }
    
    return A[L][N]
}
