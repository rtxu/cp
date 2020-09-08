func countTriplets(A []int) int {
    // i, j, k 本身是排列关系，我们只需要求出组合数，再进行排列即可
    // 如果组合形式为 3个数相等，则排列数为 1
    // 如果组合形式为 2个数相等，则排列数为 3
    // 如果组合形式为 3个数不同，则排列数为 6
    // 以 i <= j <= k 的约束枚举所有组合情况，复杂度 n^3
    // n = 1000，只要降到 n^2*logn 即可满足
    // 很好，明明就是个暴力解法的题，凭什么值 hard
    
    n := len(A)
    count := 0
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            for k := j; k < n; k++ {
                if A[i] & A[j] & A[k] == 0 {
                    if i == j && j == k {
                        count++
                    } else if i == j || j == k || i == k {
                        count += 3
                    } else {
                        count += 6
                    }
                }
            }
        }
    }
    return count
}
