/*
T(i, j) 前 i 天以 j 活动结尾的最优值
T(0, a) = T(0, b) = T(0, c) = 0
 
T(i, j) = max{T(i-1, k!=j) + happiness of j at day i}
 
*/
 
package main
 
import "fmt"
 
const N = 100005
const INF = 1e9+7
var p [N][3]int
var T [N][3]int
 
func umax(current *int, val int) {
  if val > *current {
    *current = val
  }
}
 
func max(a ...int) int {
  v := a[0]
  for i := 1; i < len(a); i++ {
    umax(&v, a[i])
  }
  return v
}
 
func main() {
  var n int
  fmt.Scanf("%d", &n)
  for i := 1; i <= n; i++ {
    fmt.Scanf("%d %d %d", &p[i][0], &p[i][1], &p[i][2])
  }
  for i := 1; i <= n; i++ {
    for j := 0; j < 2; j++ {
      T[i][j] = 0
    }
  }
  for i := 1; i <= n; i++ {
    for j := 0; j < 3; j++ {
      v := &T[i][j]
      for k := 0; k < 3; k++ {
        if j == k { continue }
        umax(v, T[i-1][k] + p[i][j])
      }
    }
  }
  
  fmt.Println(max(T[n][0], T[n][1], T[n][2]))
}
