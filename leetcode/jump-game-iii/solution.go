func canReach(arr []int, start int) bool {
    n := len(arr)
    visited := make([]bool, n)
    
    Q := make([]int, n)
    Q[0] = start
    qHead, qTail := 0, 1
    visited[start] = true
    
    checkNext := func(next int) {
        if next >= 0 && next < n && !visited[next] {
            Q[qTail] = next
            visited[next] = true
            qTail++
        }
    }
    for qHead < qTail {
        current := Q[qHead]
        if arr[current] == 0 {
            return true
        }
        qHead++
        
        checkNext(current + arr[current])
        checkNext(current - arr[current])
    }
    return false
}
