/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack []*TreeNode

func (s Stack) Empty() bool {
    return len(s) == 0
}

func (s *Stack) Push(n *TreeNode) {
    *s = append(*s, n)
}

func (s Stack) Top() *TreeNode {
    return s[len(s)-1]
}

func (s *Stack) Pop() *TreeNode {
    old := *s
    newSize := len(old)-1
    *s = old[:newSize]
    return old[newSize]
}

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    var result []int
    stack := make(Stack, 0)
    stack.Push(root)
    result = append(result, root.Val)
    var lastPop *TreeNode
    for !stack.Empty() {
        current := stack.Top()
        if current.Left != nil && lastPop == current.Left {
            if current.Right != nil {
                stack.Push(current.Right)
                result = append(result, current.Right.Val)
            } else {
                lastPop = stack.Pop()
            }
        } else if current.Right != nil && lastPop == current.Right {
            lastPop = stack.Pop()
        } else {
            if current.Left != nil {
                stack.Push(current.Left)
                result = append(result, current.Left.Val)
            } else if current.Right != nil {
                stack.Push(current.Right)
                result = append(result, current.Right.Val)
            } else {
                lastPop = stack.Pop()
            }
        }
    }
    return result
}
