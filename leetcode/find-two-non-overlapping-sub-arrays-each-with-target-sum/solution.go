// Time: O(n)
// Space: O(n)
// Sliding Window + DP

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func minSumOfLengths(arr []int, target int) int {
    n := len(arr)
    var sum int
    // A(i) 表示 0...i 满足 sum 为 target 的子数组的最小长度
    // 则最终答案为，当发现子数组 Sum[i, j] = target 时，result = min(result, j-i+1 + A[i-1])
    A := make([]int, n)
    result := n + 1
    for begin, end := 0, 0; end < n; end++ {
        sum += arr[end]
        if end > 0 {
            A[end] = A[end-1]
        }
        for sum >= target {
            if sum == target {
                if begin > 0 && A[begin-1] > 0 {
                    result = min(result, end-begin+1 + A[begin-1])
                }
                if A[end] == 0 {
                    A[end] = end-begin+1
                } else {
                    A[end] = min(A[end], end-begin+1)
                }
            }
            sum -= arr[begin]
            begin++
        } 
    }
    if result > n {
        return -1
    } else {
        return result
    }
}
