package main
 
import "fmt"
 
const N = 1e5+5
 
type Vertex struct {
  indegree int
  maxlen int
  adj []*Vertex
}
 
func (v *Vertex) add(u *Vertex) {
  v.adj = append(v.adj, u)
}
 
func (v *Vertex) incIndegree() {
  v.indegree++
}
 
func (v *Vertex) decIndegree() {
  v.indegree--
}
 
func umax(current *int, val int) {
  if val > *current {
    *current = val
  }
}
 
type Queue []*Vertex
 
func (q *Queue) Push(v *Vertex) {
  *q = append(*q, v)
}
 
func (q *Queue) Pop() *Vertex {
  old := *q
  *q = old[1:]
  return old[0]
}
 
func (q Queue) Empty() bool {
  return len(q) == 0
}
 
func main() {
  var n, m int
  fmt.Scanf("%d %d", &n, &m)
  vs := make([]*Vertex, n)
  for i := 0; i < n; i++ {
    vs[i] = &Vertex{
      indegree: 0,
      maxlen: -1,
    }
  }
  for i := 0; i < m; i++ {
    var x, y int
    fmt.Scanf("%d %d", &x, &y)
    x--
    y--
    vs[x].add(vs[y])
    vs[y].incIndegree()
  }
  q := make(Queue, 0)
  for i := 0; i < n; i++ {
    if vs[i].indegree == 0 {
      q.Push(vs[i])
    }
  }
  maxlen := 0
  for !q.Empty() {
    current := q.Pop()
    for i := 0; i < len(current.adj); i++ {
      next := current.adj[i]
      umax(&next.maxlen, current.maxlen + 1)
      umax(&maxlen, next.maxlen)
      next.decIndegree()
      if next.indegree == 0 {
        q.Push(next)
      }
    }
  }
  fmt.Println(maxlen+1)
}
