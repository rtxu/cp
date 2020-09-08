func find(x int, parent []int) int {
    if parent[x] == x {
        return x
    } else {
        parent[x] = find(parent[x], parent)
        return parent[x]
    }
}

func findCircleNum(M [][]int) int {
    n := len(M)
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    
    
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if M[i][j] == 1 {
                parent[find(j, parent)] = find(i, parent)
            }
        }
    }
    
    groupMap := map[int]bool{}
    for i := 0; i < n; i++ {
        groupMap[find(i, parent)] = true
    }
    return len(groupMap)
}
