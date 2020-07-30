/*
 
T(i, j) 面临 [i, j] 的局面，当前玩家可获得的最大分数
则最终答案 taro 获得分数，T(0, n-1)，jiro 获得分数 
如果 taro 第一步选择 a_0，T(1, n-1) 
如果 taro 第一步选择 a_n-1，T(0, n-2) 
 
T(i, j) = max{ a_i + T(afterBestChoice(i+1, j)), a_j + T(afterBestChoice(i, j-1)) }
 
其中，afterBestChoice(i, j) 为处在 [i, j] 局面的人做完最优选择后剩下的局面
 
边界：
T(i, i) = a_i
bestChoice(i, i) = i
afterBestChoice(i, i) = i+1, i
 
*/
 
package main
 
import "fmt"
 
const N = 3005
var a [N]int
var T [N][N]int
var best [N][N]int
 
func afterBestChoice(i, j int) int {
  pick := best[i][j]
  if pick == i {
    return T[i+1][j]
  } else {
    return T[i][j-1]
  }
}
 
func umax(current *int, val int, i, j, choice int) {
  if val > *current {
    *current = val
    best[i][j] = choice
  }
}
 
func main() {
  var n int
  fmt.Scanf("%d", &n)
  for i := 0; i < n; i++ {
    fmt.Scanf("%d", &a[i])
  }
  for i := 0; i < n; i++ {
    for j := i; j < n; j++ {
      T[i][j] = 0
    }
  }
  for i := 0; i < n; i++ {
    T[i][i] = a[i]
    best[i][i] = i
  }
  for i := n-1; i >= 0; i-- {
    for j := i+1 ; j < n; j++ {
      v := &T[i][j]
      umax(v, a[i] + afterBestChoice(i+1, j), i, j, i)
      umax(v, a[j] + afterBestChoice(i, j-1), i, j, j)
    }
  }
  X, Y := T[0][n-1], afterBestChoice(0, n-1)
  fmt.Println(X-Y)
}
