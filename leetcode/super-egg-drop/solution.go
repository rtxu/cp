func superEggDrop(K int, N int) int {
    // A(i, j) 表示 i 个鸡蛋走 j 步可以覆盖的层数
    // 0 <= i <= K, 0 <= j <= 最大的 j 满足 A(i, maxj-1) < 10000 && A(i, maxj) >= 10000
    // 则第一个 j 满足 A(K, j) >= N，j 即为答案
    
    A := make([][]int, K+1)
    for i := 1; i <= K; i++ {
        A[i] = make([]int, N+1)
    }
    // i = 1 时，A(1, N) = N
    for j := 1; j <= N; j++ {
        A[1][j] = j
    }
    // j = 1 时，A(i, 1) = 1
    for i := 1; i <= K; i++ {
        A[i][1] = 1
    }
    
    // A(i, j) = A[i-1][j-1] + 1 + A[i][j-1]
    for i := 2; i <= K; i++ {
        for j := 2; j <= N; j++ {
            A[i][j] = A[i-1][j-1] + 1 + A[i][j-1]
            if A[i][j] >= N {
                break
            }
        }
    }
    
    var minStepCount int
    for minStepCount = 1; minStepCount <= N; minStepCount++ {
        if A[K][minStepCount] >= N {
            break
        }
    }
    return minStepCount
}
