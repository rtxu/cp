func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func candy(ratings []int) int {
    n := len(ratings)
    if n == 0 {
        return 0
    }
    
    descendSeqStartPos := -1
    candyCnt := make([]int, n)
    candyCnt[0] = 1
    calcDescendSeqCandyCnt := func(start, end int) {
        if start != -1 {
            // [start, end] 是一个递减序列
            candyCnt[start] = max(candyCnt[start], end-start+1)
            for j := end; j > start; j-- {
                candyCnt[j] = end + 1 - j
            }
        }
    }
    // 已经给第一个人分配了一个
    for i := 1; i < n; i++ {
        if ratings[i-1] > ratings[i] {
            if descendSeqStartPos == -1 {
                descendSeqStartPos = i-1
            }
        } else if ratings[i-1] == ratings[i] {
            candyCnt[i] = 1
            calcDescendSeqCandyCnt(descendSeqStartPos, i-1)
            descendSeqStartPos = -1
        } else {
            calcDescendSeqCandyCnt(descendSeqStartPos, i-1)
            candyCnt[i] = candyCnt[i-1] + 1
            descendSeqStartPos = -1
        }
    }
    calcDescendSeqCandyCnt(descendSeqStartPos, n-1)
    
    var result int
    for i := 0; i < n; i++ {
        result += candyCnt[i]
    }
    
    return result
}
