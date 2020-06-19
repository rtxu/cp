func findMinHeightTrees(n int, edges [][]int) []int {
  degree := make(map[int]map[int]bool)
  for i := 0; i < n; i++ {
    degree[i] = make(map[int]bool)
  }
  for i := 0; i < len(edges); i++ {
    from, to := edges[i][0], edges[i][1]
    degree[from][to] = true
    degree[to][from] = true
  }
  Q := make([]int, n)
  qHead, qTail := 0, 0
  for i := 0; i < n; i++ {
    if len(degree[i]) == 1 {
      Q[qTail] = i
      qTail++
    }
  }
  leftNodeCount := n
  for leftNodeCount > 2 {
      qTailSnapshot := qTail
      for qHead < qTailSnapshot {
        current := Q[qHead]
        qHead++
        for next, _ := range degree[current] {
          delete(degree[next], current)
          if len(degree[next]) == 1 {
            Q[qTail] = next
            qTail++
          }
        }
        delete(degree, current)
        leftNodeCount--
      }
  }
  
  result := make([]int, 0, len(degree))
  for n, _ := range degree {
    result = append(result, n)
  }
  return result
}
