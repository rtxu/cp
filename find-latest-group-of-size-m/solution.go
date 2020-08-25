var parent, size, count []int

func addSet(v int) {
    parent[v] = v
    size[v] = 1
    count[1]++
}

func find(v int) int {
    if parent[v] == 0 {
        return 0
    }
    if parent[v] == v {
        return v
    } else {
        parent[v] = find(parent[v])
        return parent[v]
    }
}

func union(a, b int) {
    ra, rb := find(a), find(b)
    if ra != rb {
        count[size[ra]]--
        count[size[rb]]--
        if size[ra] > size[rb] {
            ra, rb = rb, ra
        }
        parent[ra] = rb
        size[rb] += size[ra]
        count[size[rb]]++
    }
}

func findLatestStep(arr []int, m int) int {
    n := len(arr)
    parent = make([]int, n+2)
    size = make([]int, n+2)
    count = make([]int, n+1)
    
    result := -1
    for i := 0; i < n; i++ {
        v := arr[i]
        addSet(v)
        leftParent := find(v-1)
        if leftParent != 0 {
            union(leftParent, v)
        }
        rightParent := find(v+1)
        if rightParent != 0 {
            union(rightParent, v)
        }
        if count[m] > 0 {
            result = i+1
        }
    }
    return result
}
