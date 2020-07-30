/*
T(i, j) 表示 [i, j) 的最优值，则 T(0, n) 为最终答案
 
T(i, j) = min{ T(i, k) + T(k, j) + Sum(i, j) }
*/
 
package main
 
import "fmt"
import "math"
 
const N = 405
var T [N][N]int64
var a [N]int64
var sum [N]int64
 
func rangeSum(i, j int) int64 {
  return sum[j] - sum[i]
}
 
func umin(current *int64, val int64) {
  if val < *current {
    *current = val
  }
}
 
func main() {
  var n int
  fmt.Scanf("%d", &n)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d", &a[i])
    sum[i+1] = sum[i] + a[i]
  }
  for i := 0; i < n; i++ {
    T[i][i+1] = 0
  }
  for l := 2; l <= n; l++ {
    for i := 0; i+l <= n; i++ {
      j := i+l
      T[i][j] = math.MaxInt64
      v := &T[i][j]
      for k := i+1; k < j; k++ {
        umin(v, T[i][k] + T[k][j] + rangeSum(i, j))
      }
    }
  }
  
  fmt.Println(T[0][n])
}
