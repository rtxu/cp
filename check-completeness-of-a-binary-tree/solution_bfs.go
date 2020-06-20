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
    index int
}

type Queue []QNode

func (q Queue) Empty() bool {
    return q.Size() == 0
}

func (q Queue) Size() int {
    return len(q)
}

func (q *Queue) Push(v QNode) {
    *q = append(*q, v)
}


func (q *Queue) Pop() QNode {
    old := *q
    *q = old[1:]
    return old[0]
}

func isCompleteTree(root *TreeNode) bool {
    Q := make(Queue, 0)
    Q.Push(QNode{root, 0})
    
    expectedLevelSize := 1
    for !Q.Empty() {
        currentLevelSize := Q.Size()
        for i := 0; i < currentLevelSize; i++ {
            current := Q.Pop()
            if current.index != i {
                return false
            }
            if current.tree.Left != nil {
                Q.Push(QNode{current.tree.Left, current.index*2})
            }
            if current.tree.Right != nil {
                Q.Push(QNode{current.tree.Right, current.index*2+1})
            }
        }
        if !Q.Empty() && currentLevelSize != expectedLevelSize {
            return false
        }
        expectedLevelSize <<= 1
    }
    
    return true
}
