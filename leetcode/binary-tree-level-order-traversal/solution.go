/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type QNode struct {
    tree *TreeNode
    level int
}

type Queue []QNode

func (q Queue) Empty() bool {
    return len(q) == 0
}

func (q *Queue) Push(n QNode) {
    *q = append(*q, n)
}

func (q *Queue) Pop() QNode {
    old := *q
    *q = old[1:]
    return old[0]
}

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    var result [][]int
    Q := make(Queue, 0)
    Q.Push(QNode{root, 0})
    for !Q.Empty() {
        current := Q.Pop()
        if len(result) == current.level {
            result = append(result, []int{current.tree.Val})
        } else {
            result[current.level] = append(result[current.level], current.tree.Val)
        }
        if current.tree.Left != nil {
            Q.Push(QNode{current.tree.Left, current.level+1})
        }
        if current.tree.Right != nil {
            Q.Push(QNode{current.tree.Right, current.level+1})
        }
    }
    return result
}
