/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorder(current *TreeNode) (*TreeNode, *TreeNode) {
    if current == nil {
        return nil, nil
    }
    lhead, ltail := preorder(current.Left)
    rhead, rtail := preorder(current.Right)
    current.Left = nil
    if ltail == nil && rtail == nil {
        return current, current
    } else if ltail == nil {
        current.Right = rhead
        return current, rtail
    } else if rtail == nil {
        current.Right = lhead
        return current, ltail
    } else {
        current.Right = lhead
        ltail.Right = rhead
        return current, rtail
    }
}

func flatten(root *TreeNode) *TreeNode {
    head, _ := preorder(root)
    return head
}
