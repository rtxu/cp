type point struct {
    x, y int
}

func swimInWater(grid [][]int) int {
    n := len(grid)
    deltaX := []int{0, 1, 0, -1}
    deltaY := []int{-1, 0, 1, 0}
    
    Q := make([]point, n*n)
    possible := func(h int) bool {
        if grid[0][0] <= h {
        } else {
            return false
        }
        visited := make([][]bool, n)
        for i := 0; i < n; i++ {
            visited[i] = make([]bool, n)
        }
        Q[0] = point{0, 0}
        visited[0][0] = true
        qHead, qTail := 0, 1
        
        for qTail-qHead > 0 {
            current := Q[qHead]
            qHead++

            if current.x == n-1 && current.y == n-1 {
                return true
            }

            for dir := 0; dir < 4; dir++ {
                nextX := current.x + deltaX[dir]
                nextY := current.y + deltaY[dir]
                if nextX >= 0 && nextX < n && nextY >= 0 && nextY < n && !visited[nextX][nextY] {
                    elevation := grid[nextX][nextY]
                    if elevation <= h {
                        Q[qTail] = point{nextX, nextY}
                        qTail++
                        visited[nextX][nextY] = true
                    }
                }
            }
        }
        return false
    }
    
    // left 一定不可行，right 一定可行
    left, right := 0, n*n
    for left+1 < right {
        m := (left+right)>>1
        if possible(m) {
            //fmt.Println("possible", m)
            right = m
        } else {
            //fmt.Println("impossible", m)
            left = m
        }
    }
    return right
}
