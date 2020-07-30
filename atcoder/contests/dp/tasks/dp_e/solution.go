/*
T(i, V) 前 i 个物品，获得 V 的价值所需最小容量
T(i, V) = min{T(i, V+1), T(i-1, V), T(i-1, V - v_i) + w_i }
*/
 
package main
 
import "fmt"
import "math"
 
const N = 105
const V = 1e3 * N + 5
var T [V]int
var vs, ws [N]int
 
func umin(current *int, val int) {
  if val < *current {
    *current = val
  }
}
 
func zeroOnePack(T []int, w, v int, totalW int, maxV int) {
  for i := maxV; i >= 0; i-- {
    umin(&T[i], T[i+1])
    if i - v >= 0 && T[i-v] <= totalW && T[i-v] + w <= totalW {
      umin(&T[i], T[i-v] + w)
    }
  }
}
 
func main() {
  var n, w int
  fmt.Scanf("%d %d", &n, &w)
  sumV := 0
  for i := 0; i < n; i++ {
    fmt.Scanf("%d %d", &ws[i], &vs[i])
    sumV += vs[i]
  }
  for i := 1; i <= sumV+1; i++ {
    T[i] = math.MaxInt32
  }
  T[0] = 0
  for i := 0; i < n; i++ {
    zeroOnePack(T[:sumV+2], ws[i], vs[i], w, sumV)
  }
  
  var maxV int
  for i := sumV; i >= 0; i-- {
    if T[i] < math.MaxInt32 {
      maxV = i
      break
    }
  }
  fmt.Println(maxV)
}
