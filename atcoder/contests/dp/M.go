/*
T(i, j) 前 i 个小朋友消费了 j 个糖果的方案数
则 T(n, k) 为最终答案
 
T(i, j) = Sum{T(i-1, j) + T(i-1, j-1) + ... + T(i-1, j-a[i]) } 
= rangeSum([j-a[i], j]) = rangeSum([j-a[i], j+1))
*/
package main
 
import "fmt"
 
const N = 105
const K = 1e5+5
const MOD = 1e9+7
var a [N]int
var prefixSum [K]int
 
func rangeSum (i, j int) int {
  return (prefixSum[j] - prefixSum[i] + MOD) % MOD
}
 
func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}
 
func main() {
  var n, k int
  fmt.Scanf("%d %d", &n, &k)
  for i := 1; i <= n; i++ {
    fmt.Scanf("%d", &a[i])
  }
  current := make([]int, k+1)
  current[0] = 1
  
  for i := 1; i <= n; i++ {
    for i := 0; i <= k; i++ {
      prefixSum[i+1] = (prefixSum[i] + current[i]) % MOD
    }
    next := make([]int, k+1)
    for j := 0; j <= k; j++ {
      next[j] = rangeSum(max(0, j - a[i]), j+1)
    }
    current = next
  }
  fmt.Println(current[k])
}
