func mctFromLeafValues(arr []int) int {
    // 对于根节点来说，左边最少有一个节点，以 i 作为根节点的切分位置，j < i 左子树，j >= i 右子树
    // 则所有切分根的方式 1 <= i <= n-1, 因为 [0, i) 与 [i, n) 已确定 => max 确定 => product 确定
    // 令 T(i, j) 表示 [i, j) 范围的最优值
    // 则 T(0, n) = min(T(0, i) + T(i, n) + product([0, i), [i, n)))
    n := len(arr)
    T := make([][]int64, n)
    for i := 0; i < n; i++ {
        T[i] = make([]int64, n+1)
    }
    
    for l := 2; l <= n; l++ {
        for i := 0; i + l <= n; i++ {
            j := i + l
            v := int64(math.MaxInt64)
            for k := i+1; k < j; k++ {
                v = min(v, T[i][k] + T[k][j] + product(arr, i, k, j))
            }
            T[i][j] = v
        }
    }
    
    return int(T[0][n])
}

func max(a []int) int64 {
    result := int64(a[0])
    for i := 1; i < len(a); i++ {
        if int64(a[i]) > result {
            result = int64(a[i])
        }
    }
    return result
}

func product(arr []int, start, mid, end int) int64 {
    return max(arr[start:mid]) * max(arr[mid:end])
}

func min(a, b int64) int64 {
    if a < b {
        return a
    } else {
        return b
    }
}

