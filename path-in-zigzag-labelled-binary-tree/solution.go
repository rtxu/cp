
func flip(depth, label int) int {
  if depth == 0 {
    return 0
  }
  start := 1<<(depth-1)
  n := 1<<(depth-1)
  return start + (n - 1 - (label - start))
}

func pathInZigZagTree(label int) []int {
  // start(depth) 返回每一层的 start index
  // flip(depth, label) 返回该层 label 翻转后的 flipped_label
  
  var i int
  for i = 1; label > (1<<i)-1; i++ {}
  result := make([]int, i)
  
  for j := i; j > 0; j-- {
    result[j-1] = label
    if j % 2 == 0 {
      label = flip(j, label)
      label >>= 1
    } else {
      label >>=1
      label = flip(j-1, label)
    }
  }
  return result 
}
