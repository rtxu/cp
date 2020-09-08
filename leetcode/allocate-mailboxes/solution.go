func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func minDistance(houses []int, k int) int {
    // 第一步分堆
    // 第二步确定堆内 mailbox 位置，一旦堆的范围确定，令 h 已经排好序，h1 < h2 < h3 < h4 < h5，则显然 mailbox 放在 [h1, h5] 之间任意位置，对 h1、h5 均无影响，同理放在 h2, h4 之间任意位置对 h2, h4 无影响，则必然放在 h3 位置，即中位数位置
    
    // A(i, j) 表示 house_i...house_n-1 放置 j 个 mailbox 的最优解
    // 则 A(0, k) 为最终答案
    // A(i, j) = min(A(i+k, j-1) + Cost(i, i+k)) house_i...house_i+k-1 为一堆，house_i+k...house_n-1 分成 j-1 堆的结果
    sort.Ints(houses)
    n := len(houses)
    C := make([][]int, n)
    for i := 0; i < n; i++ {
        C[i] = make([]int, n+1)
    }
    for i := 0; i < n; i++ {
        median := houses[i]
        C[i][i+1] = 0
        for left, right := i-1, i+1; left >= 0 && right < n; left, right = left-1, right+1 {
            C[left][right+1] = C[left+1][right] + abs(houses[left]-median) + abs(houses[right]-median)
        }
        if i+2 <= n {
            C[i][i+2] = abs(houses[i+1]-median)
            for left, right := i-1, i+2; left >= 0 && right < n; left, right = left-1, right+1 {
                C[left][right+1] = C[left+1][right] + abs(houses[left]-median) + abs(houses[right]-median)
            }
        }
    }
    // 初始条件：A(i, 1) = Cost(i, n)
    A := make([][]int, n)
    for i := 0; i < n; i++ {
        A[i] = make([]int, k+1)
        A[i][1] = C[i][n]
    }
    for j := 2; j <= k; j++ {
        for i := n-j; i >= 0; i-- {
            A[i][j] = 1e9
            for k := 1; i+k < n; k++ {
                A[i][j] = min(A[i][j], A[i+k][j-1] + C[i][i+k])
            }
        }
    }
    
    return A[0][k]
}
