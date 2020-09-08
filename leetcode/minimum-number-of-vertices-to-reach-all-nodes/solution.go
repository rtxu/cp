func findSmallestSetOfVertices(n int, edges [][]int) []int {
    indegree := make([]int, n)
    E := len(edges)
    for i := 0; i < E; i++ {
        _, to := edges[i][0], edges[i][1]
        indegree[to]++
    }
    var result []int
    for i := 0; i < n; i++ {
        if indegree[i] == 0 {
            result = append(result, i)
        }
    }
    return result
}
