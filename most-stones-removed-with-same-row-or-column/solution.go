type point struct {
    x, y int
}

func makePoint(a []int) point {
    return point{a[0], a[1]}
}

func find(stone point, parent map[point]point) point {
    if parent[stone] == stone {
        return stone
    } else {
        parent[stone] = find(parent[stone], parent)
        return parent[stone]
    }
}

func removeStones(stones [][]int) int {
    n := len(stones)
    parent := make(map[point]point)
    for i := 0; i < n; i++ {
        p := makePoint(stones[i])
        parent[p] = p
    }
    
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
                parent[find(makePoint(stones[i]), parent)] = find(makePoint(stones[j]), parent)
            }
        }
    }
    
    count := make(map[point]int)
    for i := 0; i < n; i++ {
        count[find(makePoint(stones[i]), parent)]++
    }
    var result int
    for _, cnt := range count {
        result += cnt-1
    }
    return result
}
