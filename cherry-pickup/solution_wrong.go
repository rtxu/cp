func findMostCherryPath(grid [][]int) [][]int {
    n := len(grid)
    P := make([][]int, n)
    for i := 0; i < n; i++ {
        P[i] = make([]int, n)
    }
    const ILLEGAL = -2
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == -1 {
                P[i][j] = -1
            } else {
                maxP := ILLEGAL
                if i-1 >= 0 {
                    maxP = P[i-1][j]
                }
                if j-1 >= 0 {
                    if maxP == ILLEGAL || P[i][j-1] > maxP {
                        maxP = P[i][j-1]
                    }
                }
                if i == 0 && j == 0 {
                    maxP = 0
                }

                if maxP == -1 {
                    P[i][j] = -1
                } else {
                    P[i][j] = maxP + grid[i][j]
                }
            }
        }
        //fmt.Println(P[i])
    }
    //fmt.Println("P done")
    return P
}

/*
// Failed TestCase
1111000
0001000
0001001
1001000
0001000
0001000
0001111
*/

func cherryPickup(grid [][]int) int {
    // P(i, j) 表示走到 (i, j) 位置能捡到的最大樱桃数
    // P(i, j) = 
    // 1) grid[i][j] = -1, P[i][j] = -1
    // 2) 令 maxP = max(P[i-1][j], P[i][j-1]) 表示从前面步骤可以获得的最大樱桃数，
    //    如果 maxP = -1 表明前面位置不可达，则 P[i][j] = -1
    //    否则，P[i][j] = maxP + grid[i][j]
    // (0, 0) 比较特殊，此时 maxP = 0
    n := len(grid)
    cherryCount := 0
    P := findMostCherryPath(grid)
    if P[n-1][n-1] == -1 {
        return 0
    }
    cherryCount += P[n-1][n-1]
    // 清空路径上的 cherry
    i, j := n-1, n-1
    for !(i == 0 && j == 0) {
        if i-1 >= 0 && P[i-1][j] + grid[i][j] == P[i][j] {
            grid[i][j] = 0
            //fmt.Printf("clear(%d, %d)\n", i, j)
            i--
            continue
        }
        if j-1 >= 0 && P[i][j-1] + grid[i][j] == P[i][j] {
            grid[i][j] = 0
            //fmt.Printf("clear(%d, %d)\n", i, j)
            j--
            continue
        }
    }
    grid[0][0] = 0
    //fmt.Println(grid)
    
    P = findMostCherryPath(grid)
    cherryCount += P[n-1][n-1]

    return cherryCount
}
