
var M = map[int]int{
    0:0,
    1:1,
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

// 本题的复杂度分析是难点，讨论区里有个精巧的算法，非常容易分析复杂度
// ref: https://leetcode.com/problems/minimum-number-of-days-to-eat-n-oranges/discuss/794847/Polylogarithmic-solution
func minDays(n int) int {
    if v, exist := M[n]; exist {
        return v
    }
    
    M[n] = 1 + min(minDays(n/3) + n%3, minDays(n/2) + n%2)
    return M[n]
}
