// Time: O(n)
// Space: O(32)

// Key Insight: 
// result[i] = the AND result of all subarrays end with arr[i] after unique
// result[i+1] = arr[i+1] + {prev & arr[i+1], any prev in result[i]}
// due to result[i] is monotone decreasing, and the only operation to decrease is to flip 1 to 0, and there's at most 31 1 in arr[i], so result[i].size is at most 32

func chmin(a *int, b int) {
    if b < *a {
        *a = b
    }
}

func abs(n int) int { 
    if n < 0 {
        return -n 
    } else {
        return n 
    }
}

func closestToTarget(arr []int, target int) int {
    n := len(arr)
    result := int(1e9+7)
    current := make(map[int]bool)
    for i := 0; i < n; i++ {
        next := make(map[int]bool)
        for num, _ := range current {
            next[num & arr[i]] = true
        }
        next[arr[i]] = true
        for num, _ := range next {
            chmin(&result, abs(num - target))
        }
        current = next
    }
    return result
}
