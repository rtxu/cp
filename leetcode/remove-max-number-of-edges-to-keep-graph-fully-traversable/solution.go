const N = 1e5 + 5

var p [2][N]int

func find_set(v int, parent []int) int {
    if v == parent[v] {
        return v
    } else {
        parent[v] = find_set(parent[v], parent)
        return parent[v]
    }
}

func union_set(a, b int, parent []int) bool {
    pa, pb := find_set(a, parent), find_set(b, parent)
    if pa != pb {
        parent[pa] = pb
        return true
    } else {
        return false
    }
}

type ByType [][]int

func (a ByType) Len() int           { return len(a) }
func (a ByType) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByType) Less(i, j int) bool { return a[i][0] > a[j][0] }

func maxNumEdgesToRemove(n int, edges [][]int) int {
    sort.Sort(ByType(edges))
    e := len(edges)
    
    for i := 1; i <= n; i++ {
        p[0][i] = i
        p[1][i] = i
    }
    removeCnt := 0
    var mergeCnt [2]int
    for i := 0; i < e; i++ {
        t, a, b := edges[i][0], edges[i][1], edges[i][2]
        used := false
        for j := 0; j < 2; j++ {
            if t & (1 << j) > 0 {
                if union_set(a, b, p[j][:]) {
                    mergeCnt[j]++
                    used = true
                }
            }
        }
        if !used {
            removeCnt++
        }
    }
    if mergeCnt[0] != n-1 || mergeCnt[1] != n-1 {
        return -1
    } else {
        return removeCnt
    }
}
