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
  var n int
  fmt.Scanf("%d", &n)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d", &h[i])
    cost[i] = INF
  }
  cost[0] = 0
  cost[1] = cost[0] + abs(h[1] - h[0])
  for i := 2; i < n; i++ {
    cost[i] = min(cost[i-1] + abs(h[i] - h[i-1]), cost[i-2] + abs(h[i] - h[i-2]))
  }
  
  fmt.Println(cost[n-1])
}
