/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type QNode struct {
  node  	*TreeNode
  index 	int64
}

type Queue []QNode

func (q Queue) Empty() bool {
  return len(q) == 0
}

func (q Queue) Size() int {
  return len(q)
}

func (q Queue) Front() QNode {
  return q[0]
}

func (q *Queue) Push(n QNode) {
  *q = append(*q, n)
}

func (q *Queue) Pop() QNode {
  old := *q
  *q = old[1:]
  return old[0]
}


func min(a, b int64) int64 {
  if a < b {
    return a
  } else {
    return b
  }
}

func max(a, b int64) int64 {
  if a > b {
    return a
  } else {
    return b
  }
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
  // bfs 将最左节点的 index 作为 offset，去除整数 overflow 的问题
  
  Q := make(Queue, 0)
  Q.Push(QNode{root, 1})
  
  var result int
  for !Q.Empty() {
    offset := Q.Front().index
    l, r := int64(math.MaxInt64), int64(math.MinInt64)
    for size := Q.Size(); size > 0; size-- {
      current := Q.Front()
      Q.Pop()
      index := current.index - offset
      l = min(l, index)
      r = max(r, index)
      if current.node.Left != nil {
        Q.Push(QNode{current.node.Left, index*2})
      }
      if current.node.Right != nil {
        Q.Push(QNode{current.node.Right, index*2+1})
      }
    }
    if int(r-l+1) > result {
      result = int(r-l+1)
    }
  }
  
  return result
}
