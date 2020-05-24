/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    isLeftChildVisited := make(map[*TreeNode]bool)
    isRightChildVisited := make(map[*TreeNode]bool)
    Q := make([]*TreeNode, 1)
    Q[0] = root
    
    var result []int
    for len(Q) > 0 {
        current := Q[len(Q)-1]
        if !isLeftChildVisited[current] {
            if current.Left == nil {
                isLeftChildVisited[current] = true
            } else {
                Q = append(Q, current.Left)
            }
        } else if !isRightChildVisited[current] {
            if current.Right == nil {
                isRightChildVisited[current] = true
            } else {
                Q = append(Q, current.Right)
            }
        } else {
            result = append(result, current.Val)
            //fmt.Println(current.Val)
            Q = Q[:len(Q)-1]
            if len(Q) > 0 {
                parent := Q[len(Q)-1]
                if parent.Left == current {
                    isLeftChildVisited[parent] = true
                } else {
                    isRightChildVisited[parent] = true
                }
            }
            
        }
    }
    
    return result
}

