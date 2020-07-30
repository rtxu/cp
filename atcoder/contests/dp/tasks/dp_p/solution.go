package main
 
import "fmt"
 
const MOD = 1e9+7
 
type Vertex struct {
  label int
  white, black int64
  adj []*Vertex
}
 
func (v *Vertex) add(u *Vertex) {
  v.adj = append(v.adj, u)
}
 
func main() {
  var n int
  fmt.Scanf("%d", &n)
  vertexs := make([]*Vertex, n)
  for i := 0; i < n; i++ {
    vertexs[i] = &Vertex{i, 0, 0, nil}
  }
  for i := 1; i < n; i++ {
    var x, y int 
    fmt.Scanf("%d %d", &x, &y)
    x, y = x-1, y-1
    vertexs[x].add(vertexs[y])
    vertexs[y].add(vertexs[x])
  }
  postOrder(vertexs[0], nil)
  fmt.Println((vertexs[0].white + vertexs[0].black) % MOD)
}
 
func postOrder(current, parent *Vertex) {
  var allWhite, total int64
  allWhite, total = 1, 1
  childCnt := 0
  for i := 0; i < len(current.adj); i++ {
    if current.adj[i] == parent { continue }
    child := current.adj[i]
    postOrder(child, current)
    allWhite *= child.white
    total *= (child.black + child.white)
    allWhite %= MOD
    total %= MOD
    childCnt++
  }
  if childCnt == 0 {
    // no child, only parent, so leaf
    current.white, current.black = 1, 1
  } else {
    current.white = total
    current.black = allWhite
  }
}
