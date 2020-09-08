// Time: O(n)
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

func shortestSubarray(A []int, K int) int {
    // 求以 i 结尾的非空、连续子数组的最大值
    // maxSum(i) = max(maxSum(i-1) + A[i], 0)，任何负数都无效，因为负数只会增加序列长度，却不贡献值
    // 如果 maxSum(i) >= K，为了找到最短序列，找到左手边最近的 j，满足 maxSum(i) - maxSum(j) >= K，则最短序列长度 i-j
    // 考虑到我们维护了一个单调递增队列，一旦遇到 maxSum[i] - maxSum[队头] >= K，则以队头元素作为起始位置的最短序列已找到
    
    n := len(A)
    maxSum := make([]int, n+1)
    stack := make([]int, 0)
    stack = append(stack, 0)
    minLen := n+1
    for i := 1; i <= n; i++ {
        maxSum[i] = max(maxSum[i-1]+A[i-1], 0)
        for len(stack) > 0 && maxSum[i] - maxSum[stack[0]] >= K {
            minLen = min(minLen, i - stack[0])
            stack = stack[1:]
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
