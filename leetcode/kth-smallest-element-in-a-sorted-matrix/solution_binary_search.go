// Time: O(N*log(max-min))
// Space: O(1)

func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    
    countLessEqual := func(target int) (int, int, int) {
        count := 0
        r, c := n-1, 0
        lessEqual, bigger := matrix[0][0], matrix[n-1][n-1]
        for r >= 0 && c < n {
            current := matrix[r][c]
            if current <= target {
                count += r+1
                c++
                if current > lessEqual {
                    lessEqual = current
                }
            } else {
                r--
                if current < bigger {
                    bigger = current
                }
            }
        }
        return count, lessEqual, bigger
    }
    
    small, big := matrix[0][0], matrix[n-1][n-1]
    for small < big {
        mid := small + (big - small) / 2
        count, lessEqual, bigger := countLessEqual(mid)
        if count == k {
            return lessEqual
        } else if count < k {
            small = bigger
        } else if count > k {
            big = lessEqual
        }
    }
    
    return small
}
