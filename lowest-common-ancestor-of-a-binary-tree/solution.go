/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func postorder(current, p, q *TreeNode) (*TreeNode, bool, bool) {
    if current == nil {
        return nil, false, false
    }
    l_lca, l_found_p, l_found_q := postorder(current.Left, p, q)
    r_lca, r_found_p, r_found_q := postorder(current.Right, p, q)
    if l_lca != nil {
        return l_lca, true, true
    } else if r_lca != nil {
        return r_lca, true, true
    } else {
        found_p := (current.Val == p.Val || l_found_p || r_found_p)
        found_q := (current.Val == q.Val || l_found_q || r_found_q)
        if found_p && found_q {
            return current, true, true
        } else {
            return nil, found_p, found_q
        }
    }
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    lca, _, _ := postorder(root, p, q)
    return lca
}
