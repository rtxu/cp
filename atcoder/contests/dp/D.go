package main
 
import "fmt"
 
const N = 105
const W = 1e5+5
 
var vs [N]int
var ws [N]int
var T [W]int64
 
/*
T(i, j) 表示前 i 个物品用了 j 容量达到的最大值
T(i, j) = max{T(i-1, j), T(i-1, j-w_i) + v_i}
*/
 
func zeroOnePack(T []int64, w, v int, totalW int) {
  for i := totalW; i - w >= 0; i-- {
    umax(&T[i], T[i-w] + int64(v))
  }
}
 
func umax(current *int64, val int64) {
  if val > *current {
    *current = val
  }
}
 
func main() {
  var n, w int
  fmt.Scanf("%d %d", &n, &w)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d %d", &ws[i], &vs[i])
  }
  for i := 0; i <= w; i++ {
    T[i] = 0
  }
  for i := 0; i < n; i++ {
    zeroOnePack(T[:w+1], ws[i], vs[i], w)
  }
  
  fmt.Println(T[w])
}
