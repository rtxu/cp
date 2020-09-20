/*
T(i, d) 表示范围为 [1, i] 的结尾为 d 的排列数
0 <= i <= n, 1 <= d <= i
则 Sum{ T(n, d) } 为最终解

T(i, d) = 
if s[i-1] = '<'，则 d 要大于 T(i-1) 的结尾数字，Sum{ T(i-1, d') } 1 <= d' < d
if s[i-1] = '>'，则 d 要小于 T(i-1) 的结尾数字，Sum{ T(i-1, d') } d <= d' <= i-1 

考虑 T(3, x) 的合法情况，假设 seq = <><
T(3, 1) = {[2, 3, 1]}
T(3, 2) = {[1, 3, 2]}
T(3, 3) = {}

构造 T(4, x) 的过程如下，因为 seq[2] = '<'
T(4, 1) = 以 1 结尾没有任何 T(3, x) 的结尾满足 < 代表的升尾条件，故 0
T(4, 2) = T(3, 1) = 1
注：T(3, 1) 的唯一答案为：[2, 3, 1] 
= 因为新序列以 2 结尾，将原序列 >= 2 的数字+1，构成新序列 => [3, 4, 1] 
= append 2 作为结尾 => [3, 4, 1, 2]
T(4, 3) = T(3, 1) + T(3, 2)
T(4, 4) = T(3, 1) + T(3, 2) + T(3, 3)

如果 seq[2] = '>' 呢？
T(4, 1) = T(3, 1) + T(3, 2) + T(3, 3)
注意：以 1 结尾的排列做变换后为 2，故 > 1
T(4, 2) = T(3, 2) + T(3, 3)
T(4, 3) = T(3, 3)
T(4, 4) = 0


*/
package main

import "fmt"

const MOD = 1e9+7

func main() {
  var n int
  fmt.Scanf("%d", &n)
  var seq string
  fmt.Scanf("%s", &seq)
  current := make([]int64, 1+1)
  current[1] = 1
  for i := 2; i <= n; i++ {
    sum := make([]int64, i+1)
    // [1, i-1]
    for j := 1; j < i; j++ {
      sum[j+1] = sum[j] + current[j]
    }
    rangeSum := func(i, j int) int64 {
      return (sum[j] - sum[i]) % MOD
    }
    next := make([]int64, i+1)
    for j := 1; j <= i; j++ {
      if seq[i-2] == '<' {
        next[j] = rangeSum(1, j)
      } else {
        next[j] = rangeSum(j, i)
      }
    }
    current = next
  }
  var result int64
  for i := 1; i <= n; i++ {
    result += current[i]
  }
  fmt.Println(result % MOD)
}
