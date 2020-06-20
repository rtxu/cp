
func dfs(current int, E map[int][]int, visited []bool, nodeCount []int) (int, int) {
  visited[current] = true
  count, sum := 1, 0
  for i := 0; i < len(E[current]); i++ {
    to := E[current][i]
    if !visited[to] {
      childCount, childSum := dfs(to, E, visited, nodeCount)
      count += childCount
      sum += childCount + childSum
    }
  }
  nodeCount[current] = count
  return count, sum
}

func sumOfDistancesInTree(N int, edges [][]int) []int {
    E := make(map[int][]int)
    for i := 0; i < N; i++ {
        E[i] = make([]int, 0)
    }
    for e := 0; e < len(edges); e++ {
        from, to := edges[e][0], edges[e][1]
        E[from] = append(E[from], to)
        E[to] = append(E[to], from)
    }
    result := make([]int, N)

    visited := make([]bool, N)
    count := make([]int, N)
    _, result[0] = dfs(0, E, visited, count)
    
    for i := 0; i < N; i++ {
        visited[i] = false
    }
    Q := make([]int, N)
    qHead, qTail := 0, 0
    Q[qTail] = 0
    qTail++
    for qHead < qTail {
        current := Q[qHead]
        visited[current] = true
        qHead++
        for e := 0; e < len(E[current]); e++ {
            next := E[current][e]
            if !visited[next] {
                Q[qTail] = next
                qTail++
                result[next] = result[current] + (N-count[next]) - count[next]
            }
        }
    }
    
    
  	return result
}
