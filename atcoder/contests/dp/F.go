/*
T(i, j) s[:i) t[:j) 的最长公共子序列长度
T(0, 0) = 0
T(i, j) = max{T(i-1, j), T(i, j-1), T(i-1, j-1) + s[i-1] == t[j-1] ? 1 : 0}
pre(i, j) 记录状态
*/
 
package main
 
import "fmt"
 
type Node struct {
  i, j int
}
 
const N = 3005
var T [N][N]int
var pre [N][N]Node
 
func umax(current *int, val int, pre_i, pre_j int, i, j int) {
  if val > *current {
    *current = val
    pre[i][j] = Node{pre_i, pre_j}
  }
}
 
func printResult(i, j int, s string) {
  result := make([]byte, 0)
  for !(i == 0 && j == 0) {
    pre_i, pre_j := pre[i][j].i, pre[i][j].j
    if T[i][j] == T[pre_i][pre_j] + 1 {
      result = append(result, s[i-1])
    }
    i, j = pre_i, pre_j
  }
  n := len(result)
  for i := n-1; i >= 0; i-- {
    fmt.Printf("%c", result[i])
  }
  fmt.Println()
}
 
func main() {
  var s, t string
  fmt.Scanf("%s", &s)
  fmt.Scanf("%s", &t)
  n, m := len(s), len(t)
  for i := 1; i <= n; i++ {
    for j := 1; j <= m; j++ {
      T[i][j] = -1
      v := &T[i][j]
      umax(v, T[i-1][j], i-1, j, i, j)
      umax(v, T[i][j-1], i, j-1, i, j)
      delta := 0
      if s[i-1] == t[j-1] {
        delta = 1
      }
      umax(v, T[i-1][j-1] + delta, i-1, j-1, i, j)
    }
  }
  
  printResult(n, m, s)
}
