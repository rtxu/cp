const INFINITY = 1e9

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
    // D(len, j) 表示路径长度至多为 len 目的地为 j 的最短路径
    // D(K+1, dst) 为最终答案
    // 初始 D(0, src) = 0
    // 转移方程：D(len, j) = min{D(len-1, j), D(len-1, u) + w}, 对于每一条边 u, j, w
    D := make([]int, n)
    for i := 0; i < n; i++ {
        D[i] = INFINITY
    }
    D[src] = 0
    nextD := make([]int, n)
    for i := 0; i <= K; i++ {
        copy(nextD, D)
        for j := 0; j < len(flights); j++ {
            u, v, w := flights[j][0], flights[j][1], flights[j][2]
            if D[u] + w < D[v] {
                nextD[v] = min(nextD[v], D[u] + w)
            }
        }
        copy(D, nextD)
    }
    if D[dst] == INFINITY {
        return -1
    } else {
        return D[dst]
    }
}
