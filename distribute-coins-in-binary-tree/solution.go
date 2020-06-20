/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(n int) int {
  if n < 0 {
    return -n
  } else {
    return n
  }
}

func dfs(current *TreeNode, moveCnt *int) (int, int) {
  if current == nil {
    return 0, 0
  }
  lcoin, lcount := dfs(current.Left, moveCnt)
  rcoin, rcount := dfs(current.Right, moveCnt)
  coin, count := lcoin + rcoin + current.Val, lcount + rcount + 1
  // coin > count, move to parent; coin < count, move from parent
  *moveCnt += abs(coin - count)
  return coin, count
}

func distributeCoins(root *TreeNode) int {
  // 一次 move 有 send 方和 recv 方，我们要求最小 move 量，仅统计 send 方
  // 对于每一条边：切断这条边，生成两颗子树，more = 当前子树拥有的coin数-节点数 > 0, 表明该节点需要向另一个子树输送 more 个 coin
  var moveCnt int
  dfs(root, &moveCnt)
  return moveCnt
}
