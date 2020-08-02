/*
T(i, j) 投完前 i 个硬币，j 个 head 的概率
T(1, 0) = 1 - p_1
T(1, 1) = p_1
 
T(i, j) = T(i-1, j) * (1_p_j) + T(i-1, j-1) * p_j
*/
package main
 
import "fmt"
 
const N = 3005
var T [N]float64
var p [N]float64
 
func main() {
  var n int
  fmt.Scanf("%d", &n)
  for i := 1; i <= n; i++ {
    fmt.Scanf("%f", &p[i])
  }
  for i := 1; i <= n; i++ {
    T[i] = 0
  }
  T[0] = 1
  for i := 1; i <= n; i++ {
    for j := i; j >= 0; j-- {
      T[j] = T[j] * (1-p[i]) 
      if j-1 >= 0 {
        T[j] += T[j-1] * p[i]
      }
    }
  }
  var res float64
  for i := (n+1)>>1; i <= n; i++ {
    res += T[i]
  }
  fmt.Println(res)
}
