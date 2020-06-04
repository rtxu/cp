func find(x int, parent []int) int {
    if parent[x] == x {
        return x
    } else {
        parent[x] = find(parent[x], parent)
        return parent[x]
    }
}

func makeConnected(n int, connections [][]int) int {
    if len(connections) < n-1 {
        return -1
    }
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    
    for e := 0; e < len(connections); e++ {
        n1, n2 := connections[e][0], connections[e][1]
        parent[find(n1, parent)] = find(n2, parent)
    }
    
    unionMap := make(map[int]bool)
    for i := 0; i < n; i++ {
        unionMap[find(i, parent)] = true
    }
    return len(unionMap)-1
}
