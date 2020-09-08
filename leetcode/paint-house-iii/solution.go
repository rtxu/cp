const INF = 1000000000

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
    // 通过观察可以看出，假设第一堆分好 (0, i) 颜色为 j，成本为 c
    // Cost(i, j, c) [i,j) 颜色为 c 的成本
    // A(i, j, c) 表示 i....n-1 的所有 houses 分成 j 堆且 A[i] 的颜色为 c 的最小成本
    // A(0, target, any C) 为最终结果
    // A(i, j, c) = min(Cost(i, i+k, c)+A(i+k, j-1, !c))
    // 0 <= i < m, 1 <= j <= target, 0 <= c < n
    
    C := make([][][]int, m)
    for i := 0; i < m; i++ {
        C[i] = make([][]int, m+1)
        for j := 0; j <= m; j++ {
            C[i][j] = make([]int, n)
            for c := 0; c < n; c++ {
                C[i][j][c] = INF
            }
        }
    }
    paintCost := func(h, c int) int {
        if houses[h] == 0 {
            return cost[h][c]
        } else {
            if c+1 == houses[h] {
                return 0
            } else {
                return INF
            }
        }
    }
    // O(m^2*n)
    for i := 0; i < m; i++ {
        for c := 0; c < n; c++ {
            var currentCost int
            for j := i+1; j <= m; j++ {
                currentCost += paintCost(j-1, c)
                if currentCost > INF {
                    currentCost = INF
                }
                C[i][j][c] = currentCost
            }
        }
    }
    
    // O(m*target*n*m) => O(m^3*n)
    A := make([][][]int, m)
    for i := 0; i < m; i++ {
        A[i] = make([][]int, target+1)
        for j := 1; j <= target; j++ {
            A[i][j] = make([]int, n)
            for c := 0; c < n; c++ {
                A[i][j][c] = INF
            }
        }
    }
    // 初始条件：j = 1 时，A(i, 1, c) = C(i, m, c)
    for i := 0; i < m; i++ {
        for c := 0; c < n; c++ {
            A[i][1][c] = C[i][m][c]
        }
    }
    minCostNotC := func(i, j, targetC int) int {
        min := INF
        for c := 0; c < n; c++ {
            if c != targetC {
                if A[i][j][c] < min {
                    min = A[i][j][c]
                }
            }
        }
        return min
    }
    for j := 2; j <= target; j++ {
        for c := 0; c < n; c++ {
            for i := m-j; i >= 0; i-- {
                min := INF
                // [i+k, m) 要被分为 j-1 堆，即 m-(i+k) >= j-1
                for k := 1; m-(i+k) >= j-1; k++ {
                    v := C[i][i+k][c] + minCostNotC(i+k, j-1, c)
                    if v < min {
                        min = v
                    }
                }
                A[i][j][c] = min
            }
        }
    }
    min := INF
    for c := 0; c < n; c++ {
        if A[0][target][c] < min {
            min = A[0][target][c]
        }
    }
    if min >= INF {
        return -1
    } else {
        return min
    }
}
