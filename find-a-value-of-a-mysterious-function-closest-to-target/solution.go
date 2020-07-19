// Time: O(n)
// Space: O(32)
// tag: sliding-window, bit-manipulation

func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func closestToTarget(arr []int, target int) int {
    n := len(arr)
    var zeroBitCount, oneBitCount [32]int
    var current int
    minV := int(1e9+17)
    for begin, end := 0, 0; end < n; end++ {
        if end - begin == 0 {
            current = arr[end]
        } else {
            current &= arr[end]
        }
        for i := 0; i <= 31; i++ {
            if arr[end] & (1<<i) > 0 {
                oneBitCount[i]++
            } else {
                zeroBitCount[i]++
            }
        }
        minV = min(minV, abs(current - target))
        for begin <= end && current < target {
            for i := 0; i <= 31; i++ {
                if arr[begin] & (1<<i) > 0 {
                    oneBitCount[i]--
                } else {
                    zeroBitCount[i]--
                    if zeroBitCount[i] == 0 && oneBitCount[i] > 0 {
                        current += (1<<i)
                    }
                }
            }
            minV = min(minV, abs(current - target))
            begin++
        }
        minV = min(minV, abs(current - target))
    }
    return minV
}
