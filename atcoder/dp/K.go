package main
 
import "fmt"
 
const N = 105
const K = 1e5+5
var a [N]int
var T [K]int
 
func main() {
  var n, k int
  fmt.Scanf("%d %d", &n, &k)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d", &a[i])
  }
  for i := 0; i <= k; i++ {
    T[i] = 0
  }
  for i := 0; i <= k; i++ {
    for j := 0; j < n && a[j] <= i; j++ {
      T[i] |= 1 - T[i-a[j]]
    }
  }
  if T[k] == 0 {
    fmt.Println("Second")
  } else {
    fmt.Println("First")
  }
}
