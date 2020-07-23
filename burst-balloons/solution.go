func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxCoins(nums []int) int {
    // A(i, j) 为 [i, j] burst 之后的最大值
    // 则 A(0, n-1) 为最终解
    // 考虑最后 burst 的气球
    // A(0, n-1) = max{
    // 最后 burst 0，A(1, n-1) + nums[-1] * nums[0] * nums[n]
    // 最后 burst 1, A(0, 0) + A(2, n-1) + nums[-1] * nums[1] * nums[n]
    // 最后 burst 2, A(0, 1) + A(3, n-1) + nums[-1] * nums[2] * nums[n]
    // 最后 burst n-1, A(0, n-2) + nums[-1] * nums[n-1] * nums[n]
    // }
    
    // 重新定义 nums 左右两边各插入 1
    // A(i, j) 表示 (i, j) burst 之后的最大值
    // A(0, n-1) 为最终解，注意此时的 n = orig_n + 2
    // 依然考虑最后 burst 的气球
    // A(0, n-1) = max {
    // 最后 burst 1, A(0, 1) + A(1, n-1) + nums[0] * nums[1] * nums[n-1]
    // 最后 burst 2, A(0, 2) + A(2, n-1) + nums[0] * nums[2] * nums[n-1]
    // ...
    // 最后 burst n-2, A(0, n-2) + A(n-2, n-1) + nums[0] * nums[n-2] * nums[n-1]
    // }
    
    arr := make([]int, len(nums)+2)
    n := len(arr)
    arr[0], arr[n-1] = 1, 1
    copy(arr[1:], nums)
    
    A := make([][]int, n)
    for i := 0; i < n; i++ {
        A[i] = make([]int, n)
    }
    for l := 2; l < n; l++ {
        for i := 0; i + l < n; i++ {
            j := i+l
            for k := i+1; k < j; k++ {
                A[i][j] = max(A[i][j], A[i][k] + A[k][j] + arr[i] * arr[k] * arr[j])
            }
        }
    }
    
    return A[0][n-1]
}
