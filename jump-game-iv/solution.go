func minJumps(arr []int) int {
    n := len(arr)
    edge := make(map[int][]int)
    for i := 0; i < n; i++ {
        edge[arr[i]] = append(edge[arr[i]], i)
    }
    
    minSteps := make([]int, n)
    for i := 1; i < n; i++ {
        minSteps[i] = -1
    }
    Q := make([]int, n)
    Q[0] = 0
    qHead, qTail := 0, 1
    
    checkNext := func(current, next int) bool {
        if next >= 0 && next < n && minSteps[next] == -1 {
            Q[qTail] = next
            qTail++
            minSteps[next] = minSteps[current]+1
            return true
        }
        return false
    }
    for qHead < qTail {
        current := Q[qHead]
        if current == n-1 {
            return minSteps[current]
        }
        qHead++
        
        checkNext(current, current+1)
        checkNext(current, current-1)
        for e := 0; e < len(edge[arr[current]]); e++ {
            next := edge[arr[current]][e]
            checkNext(current, next)
        }
        delete(edge, arr[current])
    }
    
    return -1
}
