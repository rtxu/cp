// Time: O(E^2), E is the number of edges
// Space: O(E)

type ByWeight [][]int

func (a ByWeight) Len() int { return len(a) }
func (a ByWeight) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i][2] < a[j][2] }

func find(x int, f []int) int {
  if f[x] == x {
    return x
  } else {
    f[x] = find(f[x], f)
    return f[x]
  }
}

func MST(f []int, edges [][]int, exceptEdge map[int]bool) int {
  m := len(edges)
  minCost := 0
  for i := 0; i < m; i++ {
    if exceptEdge[i] { continue }
    n1, n2, w := edges[i][0], edges[i][1], edges[i][2]
    rn1, rn2 := find(n1, f), find(n2, f)
    if rn1 != rn2 {
      f[rn1] = rn2
      minCost += w
    }
  }
  return minCost
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
  m := len(edges)
  for i := 0; i < m; i++ {
    edges[i] = append(edges[i], i)
  }
  sort.Sort(ByWeight(edges))

  f := make([]int, n)
  for i := 0; i < n; i++ { f[i] = i }
  minCost := MST(f, edges, map[int]bool{})

  critical := map[int]bool{}
  for k := 0; k < m; k++ {
    for i := 0; i < n; i++ { f[i] = i }
    sum := MST(f, edges, map[int]bool{k: true})
    if sum != minCost {
      critical[k] = true
    }
  }

  pesudoCritical := map[int]bool{}
  for k := 0; k < m; k++ {
    if critical[k] { continue }
    for i := 0; i < n; i++ { f[i] = i }
    n1, n2, w := edges[k][0], edges[k][1], edges[k][2]
    rn1, rn2 := find(n1, f), find(n2, f)
    f[rn1] = rn2
    sum := w + MST(f, edges, map[int]bool{k: true})
    if sum == minCost {
      pesudoCritical[k] = true
    }
  }

  result := make([][]int, 2)
  for edge, _ := range critical {
    result[0] = append(result[0], edges[edge][3])
  }
  for edge, _ := range pesudoCritical {
    result[1] = append(result[1], edges[edge][3])
  }

  return result
}
