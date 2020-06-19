/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// WARNING: 测试 case 不强，容易构造出 int overflow 的 case
func dfs(current *TreeNode, level int, label int, L *[][2]int) {
  if current == nil {
    return
  }
  if len(*L) == level {
    *L = append(*L, [2]int{label, label})
  } else {
    (*L)[level][1] = label
  }
  dfs(current.Left, level+1, label*2, L)
  dfs(current.Right, level+1, label*2+1, L)
}

func widthOfBinaryTree(root *TreeNode) int {
  // 想象一颗完全二叉树，对树上的每个节点标号
 	// 则第一层：1
  // 第二层：2， 3
  // 第三层：4、5、6、7
  // left = 2*parent, right = 2*parent+1
  // 第 n 层的起点为 2^(n-1) 终点为 2^(n-1) + 2^(n-1)-1 = 2^n-1
  // 通过 dfs 遍历，可以获得每个节点的编号，逐层记录最大值和最小值
  // 则最终结果为: 最大值 - 最小值 + 1
  // 问题：容易构造出标号非常大的 case
  
  if root == nil {
    return 0
  }
  
  L := make([][2]int, 0)
  dfs(root, 0, 1, &L)
  maxWidth := 0
  for i := 0; i < len(L); i++ {
      v := L[i][1] - L[i][0] + 1
    if v > maxWidth {
      maxWidth = v
    }
  }
  return maxWidth
}
