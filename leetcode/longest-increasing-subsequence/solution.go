// Time: O(n^2)
// Space: O(n)

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func lengthOfLIS(nums []int) int {
    n := len(nums)
    LIS := make([]int, n)
    var maxL int
    for i := n-1; i >= 0; i-- {
        v := 1
        for j := i+1; j < n; j++ {
            if nums[i] < nums[j] {
                v = max(v, 1 + LIS[j])
            }
        }
        LIS[i] = v
        if LIS[i] > maxL {
            maxL = LIS[i]
        }
    }
    return maxL
}
