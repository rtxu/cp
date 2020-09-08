func find(x int, parent []int) int {
    if parent[x] == x {
        return x
    } else {
        parent[x] = find(parent[x], parent)
        return parent[x]
    }
}

func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    parent := make([]int, n+1)
    for i := 1; i <= n; i++ {
        parent[i] = i
    }
    
    for e := 0; e < n; e++ {
        i, j := edges[e][0], edges[e][1]
        if find(i, parent) == find(j, parent) {
            return edges[e]
        } else {
            parent[find(i, parent)] = find(j, parent)
        }
    }
    
    // never reach here
    return nil
}
