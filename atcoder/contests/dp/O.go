/*
T(i, S) 截止到第 i 行，状态为 S 的方案数，如果截止到 i 行 S 中没有 i 个 1，则没意义
 
T(i, S' | (1<< j) ) = Sum{ T(i-1, S') }
*/
 
package main
 
import "fmt"
 
const MOD = 1e9+7
const N = 21
var a [N][N]int
 
func main() {
  var n int
  fmt.Scanf("%d", &n)
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      fmt.Scanf("%d", &a[i][j])
    }
  }
  current := make([]int, 1<<n)
  current[0] = 1
  for i := 0; i < n; i++ {
    next := make([]int, 1<<n)
    for s := 0; s < (1<<n); s++ {
      if current[s] == 0 { continue }
      for j := 0; j < n; j++ {
        if a[i][j] == 0 { continue }
        next[s | (1<<j)] = (next[s | (1<<j)] + current[s]) % MOD
      }
    }
    current = next
  }
  fmt.Println(current[(1<<n) - 1])
}
