func minCost(s string, cost []int) int {
    n := len(s)
    result := 0
    for i := 0; i < n; {
        j := i+1
        for j < n && s[i] == s[j] { j++ }
        if j - i > 1 {
            sum := 0
            max := 0
            for k := i; k < j; k++ {
                if cost[k] > max {
                    max = cost[k]
                }
                sum += cost[k]
            }
            // fmt.Println("dedup", i, j, "cost", sum - max)
            result += sum - max
        }
        i = j
    }
    return result
}
