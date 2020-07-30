package main
 
import "fmt"
 
const N = 100005
const INF = 1e9+7
var h [N]int
var cost [N]int
 
func abs(n int) int {
  if n < 0 {
    return -n
  } else {
    return n
  }
}
 
func min(a, b int) int {
  if a < b {
    return a
  } else {
    return b
  }
}
 
func main() {
  var n, k int
  fmt.Scanf("%d %d", &n, &k)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d", &h[i])
    cost[i] = INF
  }
  cost[0] = 0
  for i := 0; i < n; i++ {
    for j := i+1; j <= i+k && j < n; j++ {
      cost[j] = min(cost[j], cost[i] + abs(h[i] - h[j]))
    }
  }
  
  fmt.Println(cost[n-1])
}
