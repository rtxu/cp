/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func count(current *TreeNode, label int, visited []bool) (int, bool) {
    if current == nil {
        return 0, true
    }
    if label < 101 {
        visited[label] = true
    } else {
        return 0, false
    }
    lcount, lvalid := count(current.Left, label*2, visited)
    if !lvalid {
        return 0, false
    }
    rcount, rvalid := count(current.Right, label*2+1, visited)
    if !rvalid {
        return 0, false
    }
    return lcount + rcount + 1, true
}

func isCompleteTree(root *TreeNode) bool {
    visited := make([]bool, 101)
    nodeCount, valid := count(root, 1, visited)
    if !valid {
        return false
    }
    consecutiveVisitedCount := 1
    for i := 2; i < 101 && visited[i]; i++ {
        consecutiveVisitedCount++
    }
    return nodeCount == consecutiveVisitedCount
}
