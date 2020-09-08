func canCross(stones []int) bool {
    // A(i, k) 表示以 k 的跳跃跳到第 i 块石头是否可行
    // 因为最多只有 1100 块石头，所以 i,k <= 1100
    // 则 A(len(stones)-1, any k) 为 true 则返回 true，否则 false
    // A(0, 0) = true
    
    n := len(stones)
    positionToIndex := make(map[int]int)
    for i := 0; i < n; i++ {
        positionToIndex[stones[i]] = i
    }
    
    A := make([][]bool, n)
    for i := 0; i < n; i++ {
        A[i] = make([]bool, n)
    }
    A[0][0] = true
    for i := 0; i < n; i++ {
        // 在第 i 块石头上，看看能往前跳到哪里
        for k := 0; k < n; k++ {
            if A[i][k] {
                for delta := k-1; delta <= k+1; delta++ {
                    if delta > 0 {
                        nextPosition := stones[i] + delta
                        if index, exist := positionToIndex[nextPosition]; exist {
                            A[index][delta] = true
                            //fmt.Println(index, delta)
                            if index == n-1 {
                                return true
                            }
                        }
                    }
                }
            }
        }
    }
    
    return false
}
