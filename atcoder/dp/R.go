package main
 
import "fmt"
 
const MOD = 1e9+7
 
type Matrix [][]int64
 
func NewMatrix(n int) Matrix {
  m := make(Matrix, n)
  for i := 0; i < n; i++ {
    m[i] = make([]int64, n)
  }
  return m
}
 
func (m Matrix) multiply(rhs Matrix) Matrix {
  n := len(m)
  result := NewMatrix(n)
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      s := int64(0)
      for k := 0; k < n; k++ {
        s = (s + m[i][k] * rhs[k][j]) % MOD
      }
      result[i][j] = s
    }
  }
  return result
}
 
func (m Matrix) dump() {
  n := len(m)
  for i := 0; i < n; i++ {
    fmt.Println(m[i])
  }
}
 
func (m Matrix) pow(k int64) Matrix {
  n := len(m)
  result := NewMatrix(n)
  for i := 0; i < n; i++ {
    result[i][i] = 1
  }
  unit := m
  cnt := 1
  for k > 0 {
    if k & 1 > 0 {
      result = result.multiply(unit)
    }
    // fmt.Println("power cnt: ", cnt)
    // unit.dump()
    cnt <<= 1
    unit = unit.multiply(unit)
    k >>= 1
  }
  return result
}
 
func main() {
  var n int
  var k int64
  fmt.Scanf("%d %d", &n, &k)
  m := NewMatrix(n)
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      fmt.Scanf("%d", &m[i][j])
    }
  }
  m = m.pow(k)
  var result int64
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      result += m[i][j]
    }
  }
  fmt.Println(result % MOD)
}
