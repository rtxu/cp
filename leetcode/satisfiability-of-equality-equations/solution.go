func find(x int, parent []int) int {
    if parent[x] == x {
        return x
    } else {
        parent[x] = find(parent[x], parent)
        return parent[x]
    }
}

func equationsPossible(equations []string) bool {
    n := len(equations)
    i, j := 0, n-1
    swap := func(i, j int) {
        equations[i], equations[j] = equations[j], equations[i]
    }
    for i <= j {
        if equations[i][1] == '!' {
            swap(i, j)
            j--
        } else {
            i++
        }
    }
    // j+1 or i 为 not equal 等式的起点
    if i == n {
        // no not equal
        return true
    }
    parent := make([]int, 26)
    for c := 0; c < 26; c++ {
        parent[c] = c
    }
    for k := 0; k < i; k++ {
        n1, n2 := int(equations[k][0] - 'a'), int(equations[k][3] - 'a')
        parent[find(n1, parent)] = find(n2, parent)
    }
    
    for k := i; k < n; k++ {
        n1, n2 := int(equations[k][0] - 'a'), int(equations[k][3] - 'a')
        if find(n1, parent) == find(n2, parent) {
            return false
        }
    }
    return true
}
