// Time: O(nlogn)
// Space: O(n)

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func lower_bound(a []int, v []int, target int) int {
    n := len(a)
    if n == 0 {
        return 0
    } else if v[a[n-1]] < target {
        return n
    } else {
        // invariant: a[right] >= target
        left, right := 0, n-1
        for left < right {
            mid := left + (right-left)>>1
            if v[a[mid]] >= target {
                right = mid
            } else {
                left = mid+1
            }
        }
        return right
    }
}

func shortestSubarray(A []int, K int) int {
    // 求以 i 结尾的非空、连续子数组的最大值
    // maxSum(i) = max(maxSum(i-1) + A[i], 0)，任何负数都无效，因为负数只会增加序列长度，却不贡献值
    // 如果 maxSum(i) >= K，为了找到最短序列，找到左手边最近的 j，满足 maxSum(i) - maxSum(j) >= K，则最短序列长度 i-j
    // 即 maxSum(j) <= maxSum(i) - K
    
    n := len(A)
    maxSum := make([]int, n+1)
    stack := make([]int, 0)
    stack = append(stack, 0)
    minLen := n+1
    for i := 1; i <= n; i++ {
        maxSum[i] = max(maxSum[i-1]+A[i-1], 0)
        if maxSum[i] >= K {
            nearest := stack[lower_bound(stack, maxSum, maxSum[i]-K+1) - 1]
            minLen = min(minLen, i-nearest)
        }
        
        // maintain monotonic increasing
        for len(stack) > 0 && maxSum[stack[len(stack)-1]] >= maxSum[i] {
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    
    if minLen == n+1 {
        return -1
    } else {
        return minLen
    }
}
