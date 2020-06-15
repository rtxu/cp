func max(a, b string) string {
    if len(a) > len(b) {
        return a
    } else if len(a) < len(b) {
        return b
    } else {
        for i := 0; i < len(a); i++ {
            if a[i] < b[i] {
                return b
            } else if (a[i] > b[i]) {
                return a
            }
        }
        return a
    }
}

func largestNumber(cost []int, target int) string {
    // A(i) 表示花费 i 画出来的最大数
    // 则 A(target) 为最终答案
    // A(i) = max(d + A(i-cost(d))), 1 <= d <= 9, A(i-cost(d)) 合法
    A := make([]string, target+1)
    for i := 1; i <= target; i++ {
        A[i] = "0"
    }
    A[0] = ""
    for i := 1; i <= target; i++ {
        for d := 1; d <= 9; d++ {
            if i-cost[d-1] >= 0 && A[i-cost[d-1]] != "0" {
                A[i] = max(A[i], fmt.Sprintf("%d%s", d, A[i-cost[d-1]]))
            }
        }
        //fmt.Println(i, A[i])
    }
    return A[target]
}
