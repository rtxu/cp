// max-dot-product-of-two-subsequences

const INFINITY = 1e9

func max(a ...int) int {
    if len(a) == 0 {
        return -INFINITY
    }
    result := a[0]
    for i := 1; i < len(a); i++ {
        if a[i] > result {
            result = a[i]
        }
    }
    return result
}


func maxDotProduct(nums1 []int, nums2 []int) int {
    // A(i, j) 表示 num1[0:i], nums2[0:j] 的最优解
    // 则 A(n1, n2) 为最终答案
    // 初始条件：
    // 1) i == 0 || j == 0, A(i, j) = -oo，因为此时没取任何 subsequence
    // 2) A(1, 1) = nums1[0] * nums2[0]
    // 转移方程：A(i, j) = max {
    // 1) A(i-1, j)
    // 2) A(i, j-1)
    // 3) max(A(i-1, j-1), 0) + nums[i-1]*nums[j-1]
    // }
    n1, n2 := len(nums1), len(nums2)
    A := make([][]int, n1+1)
    for i := 0; i <= n1; i++ {
        A[i] = make([]int, n2+1)
        for j := 0; j <= n2; j++ {
            A[i][j] = -INFINITY
        }
    }
    
    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            A[i][j] = max(A[i-1][j], A[i][j-1], max(A[i-1][j-1], 0) + nums1[i-1]*nums2[j-1])
        }
    }
    
    return A[n1][n2]
}
