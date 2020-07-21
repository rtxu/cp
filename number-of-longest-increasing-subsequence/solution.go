func findNumberOfLIS(nums []int) int {
    // A(i) 表示以 nums[i] 为结尾的最长递增序列长度
    // A(i) = max(A(j)) + 1, j < i && nums[j] < nums[i]
    // 最大长度 = max(A(i)), 0 <= i < n
    
    n := len(nums)
    A := make([]int, n)
    cnt := make([]int, n)
    
    for i := 0; i < n; i++ {
        A[i] = 1
        cnt[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                if A[j] + 1 > A[i] {
                    A[i] = A[j] + 1
                    cnt[i] = cnt[j]
                } else if A[j] + 1 == A[i] {
                    cnt[i] += cnt[j]
                }
            }
        }
    }
    
    maxLen, maxLenCnt := 0, 0
    for i := 0; i < n; i++ {
        if A[i] > maxLen {
            maxLen = A[i]
            maxLenCnt = cnt[i]
        } else if A[i] == maxLen {
            maxLenCnt += cnt[i]
        }
    }
    return maxLenCnt
}
