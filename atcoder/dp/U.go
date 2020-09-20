/*
 
mask 掩码为 1 的 bit 组成一组
s[mask] 每种 mask 代表的组获得的分数
 
d[mask] = max{d[mask - k] + s[k]}，k 为 mask 的子集
 
复杂度：O(3^n)
 
*/
 
package main
 
import "fmt"
 
const N = 16
 
func max(a, b int64) int64 {
  if a > b {
    return a
  } else {
    return b
  }
}
 
func main() {
  var n int
  fmt.Scanln(&n)
  var a [N][N]int
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      fmt.Scan(&a[i][j])
    }
  }
  s := make([]int64, 1 << n)
  for m := 0; m < 1<<n; m++ {
    for i := 0; i < n; i++ {
      for j := i+1; j < n; j++ {
        if m & (1 << i | 1 << j) == (1 << i | 1 << j) {
          s[m] += int64(a[i][j])
        }
      }
    }
  }
  d := make([]int64, 1 << n)
  for m := 0; m < 1<<n; m++ {
    for sub := m;; sub = (sub-1)&m {
      d[m] = max(d[m], d[m^sub] + s[sub])
      if sub == 0 {
        break
      }
    }
  }
  
  fmt.Println(d[1<<n-1])
}
